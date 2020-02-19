package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
	Ref    string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

/**
  结构体中小写字母开头的不会被json序列化
*/
type School struct {
	address, Name string
}
type User struct {
	WxName string `json:"wx_name"`
	// omitempty当字段为空时不序列化
	Age    int    `json:"age,omitempty"`
	School School `json:"sch"`
}

func main() {
	// Marshal是编码一个json字节数组
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear", "测试中文"},
		Ref:    "JsonRes<Data:api.User>",
	}
	buf := bytes.NewBuffer([]byte{})
	jsonEn := json.NewEncoder(buf)
	// 设置不转义html实体
	jsonEn.SetEscapeHTML(false)
	_ = jsonEn.Encode(res1D)
	fmt.Println("123", buf.String(), "878")

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("使用注解的json ", string(res2B))
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	// 解码json
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// 将接口值设置为浮点数
	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	// 将json解码到某个对象中
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println("不返回字节数组 直接返回到某个Write接口")
	enc.Encode(d)
	u0 := User{
		WxName: "buffge",
		// 24,
		School: School{
			"xx路xx号",
			"光明小学",
		},
	}
	ujson, _ := json.Marshal(u0)
	fmt.Println(string(ujson))
}
