package main

import (
	"fmt"
	"testing"
)

const PI = 3.14159
const str string = "zhangsan"

// 常量的值必须是能够在编译时就能够确定的
// 因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()
// const num = getNum()
const num = 5 % 3

// 反斜杠 \ 可以在常量表达式中作为多行的连接符使用(并不行，为什么？)
/*const f = 0.693147180559945309417232121458\
176568075500134360255254120680009*/
const f1 = 1 / 0.345678 // this is a precise reciprocal
const bill = 1e9        // float constant
const hardEight = (1 << 100) >> 97

// 并行赋值
const a, b, c = 1, "2", 3.0

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

// 常量还可以做枚举
const (
	Male    = iota // 0
	Female         // 1
	Unknown        // 2
)

// 每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）
const (
	zero = iota // 0
	one         // 1
	two         // 2
)

const d = iota + 50

func TestOne(t *testing.T) {
	fmt.Println(PI, str, num)                                           // 3.14159 zhangsan 2
	fmt.Println(f1, bill, hardEight)                                    // 2.892865614820729 1e+09 8
	fmt.Println(a, b, c, d)                                             // 1 2 3 50
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 1 2 3 4 5 6
	fmt.Println(Male, Female, Unknown)                                  // 0 1 2
	fmt.Println(zero, one, two)                                         // 0 1 2
}

func getNum() int {
	return 0
}
