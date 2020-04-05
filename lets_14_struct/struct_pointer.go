package main

import "fmt"

type person struct {
	name string
	age  int
}
type Person person

func main() {

	John := newPerson("John", 20)
	fmt.Println(John)
	John.addAge()
	fmt.Println(John)
	John.ageAdd()
	fmt.Println(John)
	//tom := Person.New("Tom", 18)
	//fmt.Println("tom:", tom)
}

//构建函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//exported method Person.New should have comment or be unexported
func (p Person) New(name string, age int) *person {
	pe := new(person)
	pe.name = name
	pe.age = age
	return pe
}

//值接收者  就是copy
//这个传入的p只是内存的副本 因为struct是值类型
func (p person) addAge() {
	p.age++
}

//指针接收者 这个是内存地址直接改
//传入指针才能修改
func (p *person) ageAdd() {
	p.age++
}
