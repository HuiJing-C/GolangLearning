package main

import (
	"fmt"
)

var a int
var b string
var (
	c float32
	d bool
	e *string
)

//驼峰命名法
var stuAge = 22

//首字母大写，全局可见
var Name = "zhangsan"

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(stuAge, Name)

	//内部定义同名变量会覆盖全局变量
	var a = 5
	fmt.Println(a)
}
