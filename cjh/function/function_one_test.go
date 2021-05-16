package function

import (
	"fmt"
	"testing"
)

// 函数的顺序无关紧要，但是为了可读性，建议main函数写在最前面，其他函数按照一定逻辑顺序进行编写
// DRY原则：Don't Repeat Yourself,执行特定任务的代码只在程序里出现一次
// 所有的函数都有参数与返回值，函数参数，返回值以及他们的类型被统称为"函数签名"
// 函数有三种类型：1 普通带名字的函数，2 匿名函数或者lambda表达式，3 方法（Methods）
// 函数的调用基本格式：pkg.FunctionName(arg1,arg2,...,argn)
func TestOne(t *testing.T) {
	// 函数被调用时，实参被复制（简单而言）,然后传递给被调用函数
	user := User{name: "AAA"}
	changeName(user)
	println(user.name) // AAA

	str := "AAA"
	change(str)
	println(str) // AAA

	// 函数还可以将其他函数作为他的参数，只要这个被调用的函数的"函数签名"（返回值个数、返回值类型、顺序）与调用函数所需求的实参是一致的
	// 例如：f1(a, b, c int) f2(a, b int) (int, int, int)，那么就可以这样使用发f1(f2(a,b))
	// 函数重载：函数名相同，参数/返回值（类型，数量，顺序）不同；Go语言不允许函数重载
	// 不支持的主要原因是函数重载需要多余的类型匹配影响性能

	// 函数是一等值，他可以赋值给变量，类似a := FunC1，并且a知道自己的函数签名，给他赋一个不同签名的函数值是不可能的
	fu := Add(1, 2)
	fu = Delete(1, 2)
	// fu = AddStr(1, 2) // cannot use AddStr(1, 2) (type string) as type int in assignment
	fmt.Printf("%d\n", fu) // -1

	// 函数值之间可以相互比较，如果他们引用的是相同的函数或者nil，则认为它们是相同的函数
	// 函数不能在其他函数中声明，不过可以使用匿名函数来破除这个限制

	// 目前Go没有泛型（generic）的概念，也就是说他不支持那种支持多种类型的函数，不过在大部分情况下可以通过接口interface，特别是空接口与类型选择与/或者通过反射reflection实现相似的功能
	// 但是使用这些技术将导致代码更为复杂，性能更为底下。所以非常注重性能的场合，最好为每一个类型单独创建一个函数，而且代码可读性更强
}

func changeName(user User) {
	user.name = "BBB"
	fmt.Printf("%s\n", user.name)
}

type User struct {
	name string
}

func change(str string) {
	str = "BBB"
}

// 函数还可以以申明的方式被使用，作为一个函数类型
// 在这里，不需要函数体{}
type FunC1 func(int, int) int
type FunC2 func(int, int) string

func Add(a, b int) int {
	return a + b
}

func Delete(a, b int) int {
	return a - b
}

func AddStr(a, b int) string {
	return fmt.Sprint(a) + fmt.Sprint(b)
}
