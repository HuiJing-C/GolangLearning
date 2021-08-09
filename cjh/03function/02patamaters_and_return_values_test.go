package function

import (
	"fmt"
	"testing"
)

// 任何一个有返回值的函数都必须以return或者panic结尾，并且每一个分支都要return语句
// go语言默认是按值传递，也就是传递参数的副本，此时修改参数，原始值不会变
// 如果希望改变参数的值，需要按引用传递。也就是在传递参数的指针，虽然此时传递的也是引用（地址）的副本
// 但是此时地址副本与原地址指向的是同一块地址，此时修改引用指向的值原值也会改变
// 在函数调用时，像切片（slice）、字典（05map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显示的指出指针）
func TestPaAndResultOne(t *testing.T) {
	// 空白符“_”用来丢弃掉一些不需要的值，如只需要add
	add, _ := AddAndDelete(3, 6)
	println(add) //

	// 改变外部变量，传递指针给函数，不仅节省内存，而且赋予了函数直接改变外部变量的能力，所以给修改的变量不需要再用return语句
	u := &User2{name: "AAA"}
	changeNameByPtr(u)
	fmt.Printf("%s\n", u.name) // BBB
}

// 命名返回值，函数返回值已经在定义函数的时候已经被命名了，此时返回值已经被赋予了其类型的零值
// 未命名返回值有多个是必须用()括起来，但是命名返回值即使只有一个也要用()括起来
// 命名返回值既可以直接用return结束，也可以像未命名函数一样，返回return a,b...,n形式的结果,或者返回明确的值
// 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂
func AddAndDelete(a, b int) (add, delete int) {
	add = a + b
	delete = a - b
	// return add, delete
	// return 0, 0
	return
}
func changeNameByPtr(user *User2) {
	user.name = "BBB"
}

type User2 struct {
	name string
}
