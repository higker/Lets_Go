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

	//回文判断
	//上海自来水来自海上
	//黄山落叶松叶落山黄

	tstr := "上海自来水来自海上"
	rstr := []rune(tstr)
	//fmt.Println(string(rstr[9]))
	endIndex := (len(rstr) - 1) //减一是因为切片的index是从0开始的
	for i := 0; i <= len(rstr); i++ {
		//比较方式就是 通过index 和最后一个end_index进行比较 首位比较
		//fmt.Println(string(rstr[end_index]))
		if string(rstr[i]) == string(rstr[endIndex]) {
			fmt.Printf("rstr[%d] = %s | rstr[%d] = %s\n", i, string(rstr[i]), endIndex, string(rstr[endIndex]))
			if i == endIndex { //这里的判断下标就是防止 运行时下标一样就说明比对完成
				break
			} else { //不一样就说明没有比对完成
				endIndex--
			}
		} else {
			fmt.Println("不是回文！！！")
			break
		}
	}

	for i, s := range tstr {
		n := len(tstr) - 1 - i
		if string(s) != string(tstr[n]) { //减一是因为切片下标是从0开始
			//减i是每一次 循环的 下标移动
			fmt.Println("不是回文")
			break
		}
	}
	//通过上面案例 实现类似于 QQ敏感词汇拦截
	// str := "zyh 我操你妈的 ~"
	// mgc := "草泥马操你妈傻逼"
	// for _, v := range []rune(str) {
	// 	if unicode.In(v, mgc) {
	// 		fmt.Println("敏感词")
	// 	}
	// }
}
