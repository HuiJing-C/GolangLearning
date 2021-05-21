package function

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestClosePgkForTest(t *testing.T) {
	// 为了追踪程序的执行位置，我们可以使用runtime.Caller(),使用闭包函数来追踪函数的执行位置
	where := func() {
		_, file, line, _ := runtime.Caller(1) // 0 1 2 3各有含义
		fmt.Printf("%s:%d\n", file, line)
	}
	where() // C:/Users/CJH/GolandProjects/GolangLearning/cjh/function/close_package_for_test.go:15
	println("========")

	// 还可以使用log
	log.SetFlags(log.Llongfile)
	log.Println("123") // C:/Users/CJH/GolandProjects/GolangLearning/cjh/function/close_package_for_test.go:22: 123

	// 或者有一种更简单的做法
	where2 := log.Print
	where2("123") // C:/Users/CJH/GolandProjects/GolangLearning/cjh/function/close_package_for_test.go:25: 123
}
