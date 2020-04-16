package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//go语言json数据转换 & 结构体字段必须能让外部包访问到
type person struct {
	Name string
	Age  int
}
type student struct {
	person
	Class    string
	Learning string
}

func main() {
	s1 := student{
		person{
			"宁宁",
			18,
		},
		"5班",
		"Golang",
	}
	fmt.Println("s1 = ", s1)
	jsonS, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(string(jsonS))
	}
	fmt.Println(string(jsonS))

	//json反序列化
	var s2 student
	str := `{"Name":"宁宁","Age":18,"Class":"5班","Learning":"Golang"}`
	json.Unmarshal([]byte(str), &s2)
	//json([]byte(str), &s2) //反序列化传入指针 go里面的方法改的是副本
	fmt.Println("s2 = ", s2)
	//HttpServer()
}

func HttpServer() {
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	s1 := student{
		person{
			"宁宁",
			18,
		},
		"5班",
		"Golang",
	}
	fmt.Println("s1 = ", s1)
	json, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(json)
	}
	fmt.Fprintln(w, string(json))
}
