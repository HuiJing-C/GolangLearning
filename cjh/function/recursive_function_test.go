package function

import (
	"fmt"
	"testing"
)

// 递归函数：函数内部调用自身
// 不同函数可以相互调用形成闭环
func TestRecursiveOne(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("fibonacci(%d) is %d\n", i, fibonacci(i))
	}

	fmt.Printf("%d id odd: %t\n", 18, odd(18))
	fmt.Printf("%d id odd: %t\n", 17, odd(17))
}

func fibonacci(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}

func odd(i int) bool {
	if i == 0 {
		return false
	}
	return even(absolute(i - 1))
}

func even(i int) bool {
	if i == 0 {
		return true
	}
	return odd(absolute(i - 1))
}

func absolute(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
