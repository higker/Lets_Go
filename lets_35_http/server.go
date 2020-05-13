package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", welcome)     //set path
	http.ListenAndServe(":8080", nil) // start http server
}
func welcome(writer http.ResponseWriter, req *http.Request) { //request 处理
	fmt.Fprintf(writer, "Hello World")
}
