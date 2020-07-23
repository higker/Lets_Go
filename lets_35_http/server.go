package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/f", http.FileServer(http.Dir("/Users/ding/Documents/GO_CODE_DEV/src/Lets_Go/lets_35_http/")))
	http.HandleFunc("/", welcome)     //set path
	http.ListenAndServe(":8080", nil) // start http server
}
func welcome(writer http.ResponseWriter, req *http.Request) { //request 处理
	buf := make([]byte, req.ContentLength)
	req.Body.Read(buf)
	fmt.Println(string(buf[:]))
	fmt.Fprintf(writer, "Hello World")
}
