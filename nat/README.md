# nat
	客户端1 client c1
	客户端2 client c2
	服务器 server s
c1 -> s
		s获取到c1的端口
3.打洞过程

双打洞客户端
（1）A请求Server。
（2）B请求Server。
（3）Server把A的IP和端口信息发给B。
（4）Server把B的IP和端口信息发给A。
（5）A利用信息给B发消息。（A信任B）
（6）B利用信息给A发消息。（B信任A）

单打洞客户端
    A 请求 S
	S 保存 A 的 端口 并且为A 分配一个域名
	S 定时 发送消息给 A
	A 定时 发送消息给 S
	B 用A分配到的域名访问 S http请求
	S 发送消息给A
	A 转发消息到本地的80端口
	A返回消息给S
	S 返回给B