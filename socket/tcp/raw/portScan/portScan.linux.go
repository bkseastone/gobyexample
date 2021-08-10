package main

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/go-ping/ping"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/routing"
)

func main() {
	ipAddr := "47.101.3.179"
	pcapReadTimeout := time.Millisecond * 10
	concurrent := 500
	router, err := routing.New()
	if err != nil {
		log.Fatal("获取路由失败:", err)
	}
	var ip net.IP
	if ip = net.ParseIP(ipAddr); ip == nil {
		log.Fatalf("不是合法ip: %q\n", ipAddr)
	} else if ip = ip.To4(); ip == nil {
		log.Fatalf("不是合法ipv4: %q\n", ipAddr)
	}
	s, err := newScanner(ip, router, pcapReadTimeout, concurrent)
	if err != nil {
		log.Printf("创建扫描器失败 %v\n", err)
	}
	defer s.close()
	if err = s.scan(); err != nil {
		log.Printf("扫描失败 %v", err)
	}
}

type scanner struct {
	// 网卡信息
	iface *net.Interface
	// 目标 网关 本地 ip
	dst, gw, src net.IP
	handle       *pcap.Handle
	opts         gopacket.SerializeOptions
	buf          gopacket.SerializeBuffer
	done         chan struct{}       // 是否已扫描完成
	jobs         chan layers.TCPPort // 需要扫描的端口
	srcPort      layers.TCPPort
	ethLayer     *layers.Ethernet
	ip4Layer     *layers.IPv4
	tcpOpts      []layers.TCPOption
	concurrent   int
	netDelay     time.Duration
}

func getPingDelay(addr string) (time.Duration, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return 0, err
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return 0, err
	}
	return pinger.Statistics().AvgRtt, nil
}
func newScanner(ip net.IP, router routing.Router, timeout time.Duration, concurrent int) (*scanner, error) {
	netDelay, err := getPingDelay(ip.String())
	if err != nil {
		return nil, err
	}
	log.Println("网络延迟:", netDelay)
	s := &scanner{
		dst: ip,
		opts: gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		},
		buf:     gopacket.NewSerializeBuffer(),
		jobs:    make(chan layers.TCPPort, concurrent),
		done:    make(chan struct{}, 1),
		srcPort: 54321,
		tcpOpts: []layers.TCPOption{
			{
				OptionType:   layers.TCPOptionKindSACKPermitted,
				OptionLength: 2,
			},
		},
		concurrent: concurrent,
		netDelay:   netDelay,
	}
	// 获取本地地址 网关地址 网卡信息
	iface, gw, src, err := router.Route(ip)
	if err != nil {
		return nil, err
	}
	log.Printf("使用%v网卡扫描 ip %v, 网关: %v, 源地址: %v \n", ip, iface.Name, gw, src)
	s.gw, s.src, s.iface = gw, src, iface
	handle, err := pcap.OpenLive(iface.Name, 65536, true, timeout)
	if err != nil {
		return nil, err
	}
	s.handle = handle
	return s, nil
}
func (s *scanner) close() {
	s.handle.Close()
}
func (s *scanner) getHwAddr() (net.HardwareAddr, error) {
	start := time.Now()
	arpDst := s.dst
	if s.gw != nil {
		arpDst = s.gw
	}
	eth := layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(s.iface.HardwareAddr),
		SourceProtAddress: []byte(s.src),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte(arpDst),
	}
	buf := gopacket.NewSerializeBuffer()
	if err := s.send(buf, &eth, &arp); err != nil {
		return nil, err
	}
	for {
		if time.Since(start) > time.Second*3 {
			return nil, errors.New("timeout getting ARP reply")
		}
		data, _, err := s.handle.ReadPacketData()
		if err == pcap.NextErrorTimeoutExpired {
			continue
		} else if err != nil {
			return nil, err
		}
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
			arp := arpLayer.(*layers.ARP)
			if net.IP(arp.SourceProtAddress).Equal(arpDst) {
				return arp.SourceHwAddress, nil
			}
		}
	}
}
func (s *scanner) scan() error {

	hwaddr, err := s.getHwAddr()
	if err != nil {
		return err
	}
	s.ethLayer = &layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       hwaddr,
		EthernetType: layers.EthernetTypeIPv4,
	}
	s.ip4Layer = &layers.IPv4{
		SrcIP:    s.src,
		DstIP:    s.dst,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}
	tcp := layers.TCP{
		SrcPort: 54321,
		DstPort: 0,
		SYN:     true,
	}
	_ = tcp.SetNetworkLayerForChecksum(s.ip4Layer)
	ipFlow := gopacket.NewFlow(layers.EndpointIPv4, s.dst, s.src)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				s.scanWorker()
				wg.Done()
			}()
		}
		wg.Wait()
		time.Sleep(time.Second)
		s.done <- struct{}{}
	}()
	go func() {
		delayDuration := s.netDelay / 10 * 22
		minDelayDuration := time.Millisecond * 50
		if delayDuration < minDelayDuration {
			delayDuration = minDelayDuration
		}
		for i := 1; i < 1<<16; i++ {
			s.jobs <- layers.TCPPort(i)
			if i%s.concurrent == 0 {
				time.Sleep(delayDuration)
			}
		}
		close(s.jobs)
	}()
	begin := time.Now()
	for {
		select {
		case <-s.done:
			log.Printf("扫描完成 ,共耗时%v\n", time.Since(begin))
			return nil
		default:
		}
		// Read in the next packet.
		data, _, err := s.handle.ReadPacketData()
		if err == pcap.NextErrorTimeoutExpired {
			continue
		} else if err != nil {
			log.Printf("error reading packet: %v", err)
			continue
		}
		// Parse the packet
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		//  解析 layer
		netLayer := packet.NetworkLayer()
		if netLayer == nil {
			continue
		}
		if netLayer.NetworkFlow() != ipFlow {
			continue
		}
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		respTcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			// 正常情况下不可能发送 所有panic
			panic("tcp layer 不正确")
		}
		/*if respTcp.DstPort != tcp.SrcPort {
			continue
		}*/
		// currRespPort = respTcp.SrcPort
		if respTcp.SYN && respTcp.ACK {
			// openIPc <- int(respTcp.SrcPort)
			log.Printf("%d is open\n", respTcp.SrcPort)
		}
	}

	return nil
}
func (s *scanner) scanWorker() {
	buf := gopacket.NewSerializeBuffer()
	tcp := layers.TCP{
		SrcPort: s.srcPort,
		DstPort: 0, // will be incremented during the scan
		SYN:     true,
		Options: s.tcpOpts,
	}
	for {
		select {
		case dstPort, ok := <-s.jobs:
			if !ok {
				return
			}
			tcp.DstPort = dstPort
			_ = tcp.SetNetworkLayerForChecksum(s.ip4Layer)
			if err := s.send(buf, s.ethLayer, s.ip4Layer, &tcp); err != nil {
				// log.Printf("error sending to port %v: %v\n", tcp.DstPort, err)
			} else {
				// 	log.Printf("success sending to port %v\n", tcp.DstPort)
			}
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
func (s *scanner) send(buf gopacket.SerializeBuffer, l ...gopacket.SerializableLayer) error {
	if err := gopacket.SerializeLayers(buf, s.opts, l...); err != nil {
		return err
	}
	return s.handle.WritePacketData(buf.Bytes())
}
