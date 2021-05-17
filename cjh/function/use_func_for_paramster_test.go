package function

import (
	"fmt"
	"testing"
)

// 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行
// 将函数作为参数的最好的例子是函数 strings.IndexFunc()：
// 该函数的签名是 func IndexFunc(s string, f func(c int) bool) int，它的返回值是在函数 f(c) 返回 true、-1 或从未返回时的索引值。
func TestFuncParameter(t *testing.T) {
	callback(2, AddInt) // The sum of 1 and 2 is: 3
}
func AddInt(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(1, y)
}
