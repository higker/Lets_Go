package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//通过bufio读写文件
const (
	size = 1024
	Path = "D:/Go_workspace/src/Lets_Go/lets_18_file/ks.txt"
)

//文件http服务器测试
func main() {

	http.HandleFunc("/ks", func(w http.ResponseWriter, r *http.Request) {
		result := make(map[string]interface{})
		result["data"] = outFile(w)
		d, _ := json.Marshal(result)
		fmt.Fprint(w, string(d))
	})
	http.HandleFunc("/video", func(writer http.ResponseWriter, request *http.Request) {

		file, err := ioutil.ReadFile("C:/Users/Deen.job/Desktop/新建文件夹/out2.mp4")
		if err != nil {
			JSON, _ := json.Marshal("查无此视频")
			fmt.Fprintf(writer, string(JSON))
			return
		} else {
			writer.Header().Set("Content-Type", "video/mp4")
		}
		writer.Write(file)
	})
	fmt.Println("")
	http.ListenAndServe(":8080", nil)
}

func outFile(writer io.Writer) []string {
	file, err := os.Open(Path)
	defer file.Close() //关闭文件防止占用资源
	if err != nil {
		fmt.Fprintf(writer, "文件打开错误~~")
	}
	//通过bufio里面的方法读取文件
	reader := bufio.NewReader(file)
	content := make([]string, 0, size*8)

	for {
		temp, err := reader.ReadString('\n')
		content = append(content, temp)
		if err == io.EOF { //文件读取完毕
			break
		}
	}
	return content
}
