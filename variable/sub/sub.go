package sub

//定义一个全局变量可以被别的包访问
var GlobalVar int

func init() {
	GlobalVar = 123
}
