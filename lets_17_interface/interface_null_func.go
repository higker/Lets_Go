package main

import "fmt"

//go语言中空接口 和断言 assign

type Auto interface {
}

type lnykCo struct {
	name string
}
type BMW struct {
	name string
}
type BYD struct {
	name string
}

func main() {
	AutoSlice := []Auto{lnykCo{"领克"}, BMW{"宝马"}, BYD{"比亚迪"}}
	assign(AutoSlice)
}
func assign(as []Auto) {
	for _, v := range as {
		switch v.(type) {
		case lnykCo:
			fmt.Println(v)
		case BMW:
			fmt.Println(v)
		case BYD:
			fmt.Println(v)
		}
	}
}
