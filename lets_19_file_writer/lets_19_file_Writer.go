package main

import (
	"fmt"
	"os"
)

const (
	filePath = "D:/Go_workspace/src/Lets_Go/lets_19_file_writer/test.txt"
)

//go语言中文件写入操作
func main() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("文件打开失败~~")
	}
	defer file.Close()
	file.Write([]byte("岳宁宁在Google云计算部门,她负责的是Google Cloud CDN产品!"))
	file.WriteString("\n而我却是给菜鸟~~sorry!")
}
