package function

import (
	"fmt"
	"testing"
)

// 以a...type的形式放在参数的最后一位,a是一个类似slice的参数，如所示数组arr,可以使用arr...来调用函数
// 变长参数可以作为对应类型的slice进行二次传递
// 如果多个参数是不同类型的，有两种方式可以设计
// 1:使用结构，参数作为结构内的字段
// 2:使用空接口interface{},该方式不仅可以传递长度未知的参数，还可以接受不同类型的参数，然后使用for range结构和switch结构对每个参数的类型进行判断

func TestVariableOne(t *testing.T) {
	fmt.Printf("%d\n", AddMore(1, 2, 3, 4, 5)) // 15
	ints := []int{3, 4, 5}
	fmt.Printf("%d\n", AddMore(1, 2, ints...)) // 15
}

func AddMore(a, b int, c ...int) int {
	ints := a + b
	for i := range c {
		ints += c[i]
	}
	AddSlice(c)
	return ints
}

func AddSlice(s []int) {
	var ints = 0
	for i := range s {
		ints += s[i]
	}
	fmt.Printf("%d\n", ints)
}

// func typecheck(.., .., values … interface{}) {
// 	for _, value := range values {
// 		switch v := value.(type) {
// 		case int: …
// 		case float: …
// 		case string: …
// 		case bool: …
// 		default: …
// 		}
// 	}
// }
