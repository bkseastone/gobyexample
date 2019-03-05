package sub2

import (
	"fmt"
	//引入sub包中的变量
	. "github.com/buffge/gobyexample/variable/sub"
)

func Test() {
	fmt.Println("全局变量GlobalVar: ", GlobalVar)
}
