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
	where() // /Users/cjh/Desktop/Go/MyGitHub/GolangLearning/cjh/03function/10close_package_for_test.go:16
	println("========")

	// 还可以使用log
	log.SetFlags(log.Llongfile)
	log.Println("123") // /Users/cjh/Desktop/Go/MyGitHub/GolangLearning/cjh/03function/10close_package_for_test.go:21: 123

	// 或者有一种更简单的做法
	where2 := log.Print
	where2("123") // /Users/cjh/Desktop/Go/MyGitHub/GolangLearning/cjh/03function/10close_package_for_test.go:25: 123
}
