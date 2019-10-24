package main

import (
	"fmt"
	"math"
)

//go语言的 （float）浮点数
func main() {
	//在go语言中的2小数默认类型就是float64类型
	//如果需要使用float32就使用显示声明
	fmt.Printf("float64最大值%v\n", math.MaxFloat64)
	fmt.Printf("float32最小值%v\n", math.MaxFloat32)
	f1 := float32(0.32)
	fmt.Printf("f1 value = %v 类型是:%T", f1, f1)
}

// float64最大值1.7976931348623157e+308
// float32最小值3.4028234663852886e+38
// f1 value = 0.32 类型是:float32
