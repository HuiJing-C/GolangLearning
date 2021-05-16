package control_struct

import (
	"fmt"
	"testing"
)

func TestBreakAndContinue(t *testing.T) {
	// 对于for循环，break跳出的是最内部的循环体，而对于switch&select语句，跳出的是整个代码块，执行后续操作
	i := 3
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i <= 0 {
			break
		}
	}

	// continue是忽略本次循环，进入下一次循环(只能用于for循环)。单并不是无条件执行下次循环，执行前依然要满足循环条件
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
	// 0 1 2 3 4 6 7 8 9
}
