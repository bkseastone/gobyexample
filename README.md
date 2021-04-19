# gobyexample
study from https://gobyexample.com/

## go build 选项
go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

-w 去掉DWARF调试信息，得到的程序就不能用gdb调试了。

-s 去掉符号表,panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果，

-X 设置包中的变量值

要做的事

端口扫描
syn http slow

去除字节数组中所有空值

go mod tidy: 格式化模块,删除无用的,下载有用的

test gitpod

## 调试 init函数
设置GODEBUG = inittrace=1 会打印所有init函数调用信息



