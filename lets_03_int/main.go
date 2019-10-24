package main

import (
	"fmt"
	"math"
)

//https://www.liwenzhou.com/posts/Go/02_datatype/
func main() {
	fmt.Println("有符号的int8范围是", math.MinInt8, "到", math.MaxInt8)
	fmt.Println("有符号的int16范围是", math.MinInt16, "到", math.MaxInt16)
	fmt.Println("有符号的int32范围是", math.MinInt32, "到", math.MaxInt32)
	fmt.Println("有符号的int64范围是", math.MinInt64, "到", math.MaxInt64)
	fmt.Println("------------------------------------------------------")
	fmt.Println("无符号的uint8范围是0到", math.MaxUint8)
	fmt.Println("无符号的uint16范围是0到", math.MaxUint16)
	fmt.Println("无符号的uint32范围是0到", math.MaxUint32)
	fmt.Println("无符号的uint64范围是0到", uint64(math.MaxUint64))
}
