package main

import "github.com/google/gopacket/pcap"

// 过程
/*
c syn
s syn ack(1)
c ack(1)
c client hello 517
s ack(518)
s server hello 1400
s seg1 1400
s seg2 1296
c ack(4097)
s segEnd 328 公钥 选择的加密方式 hello done
c ack(4425)
c 服务端与客户端互换算法需要的参数 93
c 加密数据1 99
c 加密数据2 443
s ack(611 = 518 + 93)
s ack(710 = 611 + 99)
s ack(1153 = 710 + 443)
s new session ticket,change cipher spec,加密数据 274
s 加密数据 78
c ack(4777 = 4425+ 274+78)
c 加密数据(实际请求) 38
s ack(1191= 1153+38)
s 加密数据(实际返回) 800
c ack(5577 = 4777+800)
*/
/*


 */
type (
	tlsSimulator struct {
		h pcap.Handle
	}
)

func newTlsSimulator() {

}
func main() {

}
