package struct_and_method

import (
	"fmt"
	"testing"
)

// 结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。
// 匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。

// 可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。
// Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS // anonymous field
}

func TestInnerStruct(t *testing.T) {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 is: %+v\n", outer2)
}

// 通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。
// 同样地结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用，如同上面例子中那样。外层结构体通过outer.in1 直接进入内层结构体的字段
// 内嵌结构体甚至可以来自其他包。内层结构体被简单的插入或者内嵌进外层结构体。这个简单的“继承”机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func TestInnerStruct2(t *testing.T) {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

// 命名冲突
// 当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
// 外层名字会覆盖内层名字，这提供了一种重载字段或方法的方式
// 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。
// 没有办法来解决这种问题引起的二义性，必须由程序员自己修正。

type AA struct {
	a, b int
}

type BB struct {
	AA
	b string
}

func TestInnerStruct3(t *testing.T) {
	bb := BB{AA{1, 2}, "b"}
	fmt.Printf("%s\n", bb.b)    // b
	fmt.Printf("%d\n", bb.AA.b) // 2
}
