package string

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFive(t *testing.T) {
	// strconv.IntSize查看本系统平台下int所占的位数
	println(strconv.IntSize)

	//数字转字符串
	println(strconv.Itoa(123456))
	println(strconv.FormatFloat(123.567, 101, 2, 64)) //格式b e（101） f g,精度2,64=float64,32=float32

	atoi, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	println(atoi) //123
	float, err2 := strconv.ParseFloat("123.678", 64)
	if err2 != nil {
		fmt.Printf("%v\n", err)
	}
	print(float)
}
