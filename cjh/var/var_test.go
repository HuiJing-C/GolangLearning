package _var

import (
	"fmt"
	"testing"
)

var a int
var b string
var (
	c float32
	d bool
	e *string
)

// 驼峰命名法
var stuAge = 22

// 首字母大写，全局可见
var Name = "zhangsan"

func TestOne(t *testing.T) {
	fmt.Println(a, b, c, d, e) // 0  0 false <nil>
	fmt.Println(stuAge, Name)  // 22 zhangsan

	// 内部定义同名变量会覆盖全局变量
	var a = 5
	fmt.Println(a) // 5
}
