package error

import (
	"fmt"
	"github.com/HuiJing-C/GolangLearning/cjh/10error/parse"
	"testing"
)

func TestUDOne(t *testing.T) {
	// 最佳实践
	/*
		1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
		2）向包的调用者返回错误值（而不是 panic）。
		在包内部，特别是在非导出函数中有很深层次的嵌套调用时，对主调函数来说用 panic 来表示应该被翻译成错误的错误场景是很有用的（并且提高了代码可读性）
	*/
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err) // 自定义错误类型，有String()没Error()此处打印String(), 有Error()此处打印Error()
			continue
		}
		fmt.Println(nums)
	}
}

func TestUDTwo(t *testing.T) {
	var e interface{}
	e = &UDError{s: "666"}
	// String()&Error()
	// println(e) // (0x1180f58,0x12656f8)
	// fmt.Println(e) // Error()

	// String()
	// println(e) // (0x112df80,0x12656f8)
	// fmt.Println(e) // String()

	// Error()
	// println(e) // (0x112df80,0x12656f8)
	// fmt.Println(e) // Error()

	//
	println(e)     // (0x1122c40,0xc00004c5c0)
	fmt.Println(e) // &{666}

	// 自定义error使用fmt.Println()打印时 有Error()打印Error(),没Error()打印String()。什么都没有，打印结构体
}

type UDError struct {
	s string
}

// func (u *UDError) String() string {
// 	return "String()"
// }

// func (u *UDError) Error() string {
// 	return "Error()"
// }
