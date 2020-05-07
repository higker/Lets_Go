// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/5 - 6:07 下午

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "日本 韩国 123abcDF 3.1415 21 1999.08.09 22.22"
	////解析正则表达式，如果成功返回解释器
	compile := regexp.MustCompile(`\d+\.\d+\.\d+`)
	//解释失败，返回nil
	if compile == nil {
		fmt.Println("regexp is error.")
		return
	}
	//根据规则提取关键信息
	subMatch := compile.FindAllStringSubmatch(str, -1)
	fmt.Println(subMatch)
}
