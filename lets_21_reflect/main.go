package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	var p person
	str := `{"name":"john","age":19}`
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p)

	fmt.Println(reflect.TypeOf(p))        //通过反射获取类型
	fmt.Println(reflect.TypeOf(p).Kind()) // 通过反射获取底层类型 因为person是通过struct来实现的使用底层还是struct
	fmt.Println(reflect.TypeOf(p).Name()) // 通过反射获取底类型名称 不是带带package名字的
	fmt.Println(reflect.ValueOf(p))       //通过反射获取到类型里面存储的值
}
