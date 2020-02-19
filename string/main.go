package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

var (
	pl = fmt.Println
	pf = fmt.Printf
)

type point struct {
	x, y int
}

func main() {
	// typeName := "aaa<buffge>"
	pf("%q", s.Fields(" sss ")[0])
	os.Exit(1)
	pl("字符串中是否包含某值:  ", s.Contains("test", "es"))
	pl("字符串中子串的个数:     ", s.Count("test", "t"))
	pl("检查字符串是否以什么字符串开头: ", s.HasPrefix("test", "te"))
	pl("检查字符串是否以什么字符串结尾: ", s.HasSuffix("test", "st"))
	pl("检测子字符串在字符串中的index:     ", s.Index("test", "e"))
	// 分隔符可以不要
	pl("将字符串数组组合为字符串:      ", s.Join([]string{"a", "b"}, "-"))
	pl("将字符串重复输出n次:    ", s.Repeat("a", 5))
	pl("将字符串中的某值全部换为某值:   ", s.Replace("foo", "o", "0", -1))
	pl("将字符串中的某值换为某值一次:   ", s.Replace("foo", "o", "0", 1))
	pl("将字符串按照分隔符切分为字符串数组:     ", s.Split("a-b-c-d-e", "-"))
	pl("将字符串全部转换为小写:   ", s.ToLower("TEST"))
	pl("将字符串全部转换为大写:   ", s.ToUpper("test"))
	pl()
	pl("字符串长度: ", len("hello"))
	pl("取字符串中的某个字符:", string("hello"[1]))
	p0 := point{1, 2}
	pf("显示正常的值 应该是String接口 %v\n", p0)
	pf("如果v是结构体+v将会显示键值 %+v\n", p0)
	pf("打印值的语法表示 %#v\n", p0)
	pf("打印值的类型 %T\n", p0)
	pf("布尔值 %t\n", true)
	pf("十进制数字 %d\n", 123)
	pf("二进制数字 %b\n", 14)
	pf("字符 %c\n", 33)
	pf("十六进制数字 %x\n", 456)
	pf("浮点数 %f\n", 78.9)
	pf("以科学表达式显示数字 %e\n", 123400000.0)
	pf("同上 大写E %E\n", 123400000.0)
	pf("不带双引号的字符串 %s\n", "\"string\"")
	pf("带双引号的字符串 %q\n", "\"string\"")
	pf("十六进制字符串 %x\n", "hex this")
	pf("指针 %p\n", &p0)
	// 6d 表示 占6个位置
	pf("|%6d|%6d|\n", 12, 345)
	// .2表示浮点数保存2位小数
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// -表示左对齐
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	pf("|%6s|%6s|\n", "foo", "b")
	pf("|%-6s|%-6s|\n", "foo", "b")
	str := fmt.Sprintf("将字符串保存到字符串变量中 a %s", "string")
	pl(str)
	fmt.Fprintf(os.Stderr, "将字符串格式化保存到有输入接口的变量中 an %s\n", "error")
	pl([]byte("字符串转字节数组"))
	bs := make([]byte, 5)
	for k := range bs {
		bs[k] = 'n'
	}
	pl("字符数组转字符串 ", string(bs))
	pl("数字转字符串 ", strconv.Itoa(888))
	i2, _ := strconv.Atoi("34234")
	pl("字符串转数字", i2)
	name := "buffge"
	// 二进制转16进制字符串
	nameHex, _ := Bin2hex(name)
	fmt.Printf("二进制 %s 转16进制 hex: %v\n", name, nameHex)
	// 二进制转16进制字节数组
	nameHexBts, _ := Bin2hexBytes(name)
	fmt.Printf("二进制 %s 转16进制 hex bytes: %v\n", name, nameHexBts)

	// 16进制字符串|[]byte转二进制字符串
	nameBin, _ := Hex2Bin(nameHex)
	fmt.Printf("十六进制字符串 %v 转二进制 bin: %v\n", nameHex, nameBin)
	nameBin2, _ := Hex2Bin(nameHexBts)
	fmt.Printf("十六进制bytes %v 转二进制 bin: %v\n", nameHexBts, nameBin2)
	nameBin3, _ := Hex2BinBytes(nameHex)
	fmt.Printf("十六进制字符串 %v 转二进制字符数组 bin: %v\n", nameHex, nameBin3)
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 字符串转 int
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` 会自动识别出十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// `ParseUint` 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// `Atoi` 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
}
func Bin2hex(bin interface{}) (hexStr string, err error) {
	var bts []byte
	switch bin.(type) {
	case string:
		bts = []byte(bin.(string))
	case []byte:
		bts = bin.([]byte)
	default:
		err = errors.New("require string or []byte")
	}
	// 二进制转16进制字符串
	hexStr = hex.EncodeToString(bts)

	return
}
func Bin2hexBytes(bin interface{}) ([]byte, error) {
	hexStr, err := Bin2hex(bin)
	return []byte(hexStr), err
}

// 16进制字符串|[]byte转二进制字符串
func Hex2Bin(hexData interface{}) (string, error) {
	var (
		hexBys []byte
		err    error
	)
	switch hexData.(type) {
	case string:
		hexBys = []byte(hexData.(string))
	case []byte:
		hexBys = hexData.([]byte)
	default:
		err = errors.New("require string or []byte")
	}
	maxEnLen := hex.EncodedLen(len(hexBys))
	binBys := make([]byte, maxEnLen)
	hex.Decode(binBys, hexBys)
	return string(binBys), err
}
func Hex2BinBytes(bin interface{}) ([]byte, error) {
	binStr, err := Hex2Bin(bin)
	return []byte(binStr), err
}
