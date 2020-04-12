package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//go implement simple linux cat command
//cat D:/Go_workspace/src/Lets_Go/lets_19_file_writer/cat_main.go
func main() {
	args := os.Args
	//fmt.Println(args)
	if cap(args) > 1 {
		fmt.Println(cat(args[1]))
	} else {
		fmt.Println("你传入源文件地址！！！")
	}
}

func cat(src string) string {
	con, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("文件读取错误~")
	}
	return string(con)
}
