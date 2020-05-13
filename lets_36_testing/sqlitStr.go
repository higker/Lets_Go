package lets_36_testing

import "fmt"

// 切割字符串
func Sqlit(src, sep string) []string {
	temp := make([]string, 0, 10)
	srcTemp := []rune(src)
	for _, v := range srcTemp {
		t := string(v)
		fmt.Printf("%T %T\n", sep, t)
		if !(sep == t) {
			temp = append(temp, string(v))
		}
	}
	fmt.Println(temp)
	return temp
}
