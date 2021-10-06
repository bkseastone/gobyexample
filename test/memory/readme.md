# benchmark

> go test -bench="." -v  -benchmem  -gcflags="-N -l -m=2"
> 
> -gcflags=-l 禁止内联
>
> -gcflags=-l=2 内联级别2,更积极，可能更快，可能会制作更大的二进制文件
>
> -gcflags=-l=3 内联级别3,再次更加激进，二进制文件肯定更大，也许更快，但也许会有 bug
>
> -gcflags=-l=4 内联级别4,在 Go 1.11 中将支持实验性的 中间栈内联优化。
>
> -gcflags=-N 禁止优化代码
> -gcflags=-m 显示优化决策