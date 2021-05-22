package function

import (
	"fmt"
	"testing"
	"time"
)

// 程序的运行时间是校验程序性能的基本标准，可以使用time下的Now()与Sub()来计算
func TestRunTime(t *testing.T) {
	now := time.Now()
	time.Sleep(time.Second * 3)
	end := time.Now()
	delta := end.Sub(now)
	fmt.Printf("The progress run %s\n", delta)
}

// 还可以通过内存缓存来提高性能，比如斐波那契函数，前面的值其实已经计算过了，当计算后面的值时，不需要再重复计算前面的值了
var fbs [50]uint64

func TestOptimizeFibonacci(t *testing.T) {
	now := time.Now()
	for i := 0; i < 40; i++ {
		fmt.Printf("fibonacci2(%d) is %d\n", i, fibonacci2(i))
	}
	end := time.Now()
	duration := end.Sub(now)
	fmt.Printf("Progress running %s\n", duration) // Progress running 90.333µs 对比recursive_function_test.go中的1.138394463s这个值会随着40加大而更加明显
	// 内存缓存的技术在使用计算成本相对昂贵的函数时非常有用（不仅限于例子中的递归），譬如大量进行相同参数的运算。这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数。
}

func fibonacci2(n int) (res uint64) {
	if fbs[n] != 0 {
		res = fbs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fbs[n] = res
	return
}
