package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API = "http://www.ibyte.me"
)

func main() {
	response, _ := http.Get(API)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
