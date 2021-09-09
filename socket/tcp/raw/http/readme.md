# raw socket 实现http

丢弃内核发送的rst包
iptables -t filter -I OUTPUT -p tcp -d 11.1.11.1 --dport 80 --tcp-flags RST RST -j DROP