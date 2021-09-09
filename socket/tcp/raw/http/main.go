package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/routing"
)

type (
	httpSimulator struct {
		h            *pcap.Handle
		httpUrl      *url.URL
		iface        *net.Interface
		dst, gw, src net.IP
		opts         gopacket.SerializeOptions
		buf          gopacket.SerializeBuffer
		srcPort      layers.TCPPort
		dstPort      layers.TCPPort
		ethLayer     *layers.Ethernet
		ip4Layer     *layers.IPv4
		tcpOpts      []layers.TCPOption
		ipFlow       gopacket.Flow
		lastSeq      uint32
	}
)

func newHttpSimulator(reqUrl string, router routing.Router) (*httpSimulator, error) {
	httpUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}
	ipAddr, err := net.ResolveIPAddr("ip4", httpUrl.Hostname())
	if err != nil {
		return nil, err
	}
	serverPort := httpUrl.Port()
	dstPort := 80
	if serverPort != "" {
		dstPort, _ = strconv.Atoi(serverPort)
	}
	s := &httpSimulator{
		dst:     ipAddr.IP,
		httpUrl: httpUrl,
		opts: gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		},
		buf: gopacket.NewSerializeBuffer(),
		tcpOpts: []layers.TCPOption{
			{
				OptionType:   layers.TCPOptionKindSACKPermitted,
				OptionLength: 2,
			},
		},
		srcPort: 54326,
		dstPort: layers.TCPPort(dstPort),
	}
	iface, gw, src, err := router.Route(ipAddr.IP)
	if err != nil {
		return nil, err
	}
	log.Printf("使用%v网卡请求 ip %v, 网关: %v, 源地址: %v \n", iface.Name, ipAddr.IP, gw, src)
	s.gw, s.src, s.iface = gw, src, iface
	h, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	s.h = h
	return s, nil
}
func main() {
	reqUrl := "http://buffge.com:65535"
	router, err := routing.New()
	if err != nil {
		log.Fatal("获取路由失败:", err)
	}
	s, err := newHttpSimulator(reqUrl, router)
	if err != nil {
		log.Printf("创建扫描器失败 %v\n", err)
	}
	defer s.close()
	if err = s.run(); err != nil {
		log.Println("run err", err)
	}
}
func (s *httpSimulator) run() error {
	hwAddr, err := s.getHwAddr()
	if err != nil {
		return err
	}
	s.ethLayer = &layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       hwAddr,
		EthernetType: layers.EthernetTypeIPv4,
	}
	s.ip4Layer = &layers.IPv4{
		SrcIP:    s.src,
		DstIP:    s.dst,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}
	s.ipFlow = gopacket.NewFlow(layers.EndpointIPv4, s.dst, s.src)
	// syn
	s.sendSyn()
	// read syn ack
	if err = s.readSynAck(); err != nil {
		return errors.New("连接到服务器失败" + err.Error())
	}
	// ack
	s.sendSynAck()
	// req
	s.sendReq()
	// time.Sleep(time.Hour)
	// read resp
	if err = s.readResp(); err != nil {
		return errors.New("服务器响应错误" + err.Error())
	}
	// fin
	// read ack || ack and fin
	// [ read fin ]
	// ack
	return nil

}
func (s *httpSimulator) sendSyn() {
	serverPort := s.httpUrl.Port()
	dstPort := 80
	if serverPort != "" {
		dstPort, _ = strconv.Atoi(serverPort)
	}
	tcp := &layers.TCP{
		SrcPort: s.srcPort,
		DstPort: layers.TCPPort(dstPort),
		SYN:     true,
	}
	_ = tcp.SetNetworkLayerForChecksum(s.ip4Layer)
	_ = s.send(s.buf, s.ethLayer, s.ip4Layer, tcp)
	log.Println("发送syn")
}
func (s *httpSimulator) close() {
	s.h.Close()
}
func (s *httpSimulator) getHwAddr() (net.HardwareAddr, error) {
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
		data, _, err := s.h.ReadPacketData()
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
func (s *httpSimulator) send(buf gopacket.SerializeBuffer, l ...gopacket.SerializableLayer) error {
	if err := gopacket.SerializeLayers(buf, s.opts, l...); err != nil {
		return err
	}
	return s.h.WritePacketData(buf.Bytes())
}

func (s *httpSimulator) readSynAck() error {
	begin := time.Now()
	for {
		if time.Since(begin) > time.Second {
			return errors.New("建立连接失败")
		}
		// Read in the next packet.
		data, _, err := s.h.ReadPacketData()
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
		if netLayer.NetworkFlow() != s.ipFlow {
			continue
		}
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		respTcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			panic("tcp layer 不正确")
		}
		if respTcp.DstPort != s.srcPort {
			continue
		}
		// currRespPort = respTcp.SrcPort
		if respTcp.SYN && respTcp.ACK {
			s.lastSeq = respTcp.Seq
			log.Println("获取 syn ack 成功")
			return nil
		}
	}
	return nil
}
func (s *httpSimulator) sendSynAck() {
	tcp := &layers.TCP{
		SrcPort: s.srcPort,
		DstPort: s.dstPort,
		ACK:     true,
		Ack:     s.lastSeq + 1,
		Seq:     1,
		Window:  54396,
	}
	_ = tcp.SetNetworkLayerForChecksum(s.ip4Layer)
	_ = s.send(s.buf, s.ethLayer, s.ip4Layer, tcp)
	log.Println("发送 syn ack 成功")
}

func (s *httpSimulator) sendReq() {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("GET %s HTTP/1.1\r\n"+
		"Host: %s\r\n"+
		"Connection: keep-alive\r\n"+
		"\r\n", s.httpUrl.RequestURI(), s.httpUrl.Host))
	httpPayload := buf.Bytes()
	tcp := &layers.TCP{
		SrcPort: s.srcPort,
		DstPort: s.dstPort,
		ACK:     true,
		Ack:     s.lastSeq + 1,
		Seq:     1,
		Window:  54396,
		PSH:     true,
		Padding: httpPayload,
	}
	_ = tcp.SetNetworkLayerForChecksum(s.ip4Layer)
	_ = s.send(s.buf, s.ethLayer, s.ip4Layer, tcp)
	log.Println("发送 req 成功")
}

func (s *httpSimulator) readResp() error {
	begin := time.Now()
	for {
		if time.Since(begin) > time.Second {
			return errors.New("建立连接失败")
		}
		// Read in the next packet.
		data, _, err := s.h.ReadPacketData()
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
		if netLayer.NetworkFlow() != s.ipFlow {
			continue
		}
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		respTcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			panic("tcp layer 不正确")
		}
		if respTcp.DstPort != s.srcPort {
			continue
		}
		if len(respTcp.Payload) > 0 {
			bts := respTcp.LayerPayload()
			log.Println("resp: ", bytes2Str(bts))
			return nil
		}
	}
	return nil
}
func str2Bytes(str string) []byte {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	header.Len = len(str)
	header.Cap = header.Len
	return *(*[]byte)(unsafe.Pointer(header))
}
func bytes2Str(bts []byte) string {
	return *(*string)(unsafe.Pointer(&bts))
}
