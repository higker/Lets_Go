package lets_36_testing

import (
	"reflect"
)

// 切割字符串
func Sqlit(src, sep string) []rune {
	var temp []rune
	for _, v := range []rune(src) {
		//fmt.Printf( "%T %T\n", sep, t)
		if !reflect.DeepEqual(v, []rune(sep)) {
			temp = append(temp, v)
		}
	}
	//fmt.Println(temp)
	return temp
}
