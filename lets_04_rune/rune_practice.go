package main

import (
	"fmt"
	"strings"
	"unicode"
)

//练习题 判断一个字符串里面有多少个中文
func main() {
	var tempStr string = "今天我学习了Go语言了."
	count := 0
	for _, v := range tempStr {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println("count = ", count)

	//计算字符串中的每个单词出现的次数
	wordStr := "How do you do"              //声明一个英文字符串遍历
	sliceStr := strings.Split(wordStr, " ") //通过split函数切割字符串
	cmap := make(map[string]int, 20)        //声明一个map存储 记录英文单词出现的次数
	for _, v := range sliceStr {
		if _, not := cmap[v]; !not {
			cmap[v] = 1
		} else {
			cmap[v]++
		}
	}
	fmt.Println(cmap)

	//通过上面案例 实现类似于 QQ敏感词汇拦截
	str := "zyh 我操你妈的 ~"
	runeStr := []rune(str)
	for _, v := range runeStr {
		t := fmt.Sprintf("%s", v);
		if  t == string("操") {
			fmt.Println("保护敏感词", t)
		}
		fmt.Println("保护敏感词", t)
	}
}
