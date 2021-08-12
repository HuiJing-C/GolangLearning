package _8interface_reflect

import (
	"fmt"
	"testing"
)

// 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
type Namer interface {
	firstName() string
}

type AmericaNamer struct{}
type Chinese struct{}

// 接口被隐式地实现。多个类型可以实现同一个接口
func (a *AmericaNamer) firstName() string {
	return "Tom"
}

func (a *AmericaNamer) Name() string {
	return "Black.Tom"
}

func (c *Chinese) firstName() string {
	return "曹"
}

// 不像大多数面向对象编程语言，在 Go 语言中接口可以有值，一个接口类型的变量或一个 接口值 ：
// var ai Namer，ai 是一个多字（multiword）数据结构，它的值是 nil。它本质上是一个指针，虽然不完全是一回事。指向接口值的指针是非法的，它们不仅一点用也没有，还会导致代码错误

func TestInterface(t *testing.T) {
	var ai Namer
	fmt.Printf("%v\n", ai)  // <nil>
	fmt.Printf("%p\n", &ai) // 0xc0000964f0
	// println(*ai) // invalid indirect of ai (type Namer)

	// 多态
	ai = &AmericaNamer{}               // 去掉&就错了，因为ai本质上是一个指针
	fmt.Printf("%s\n", ai.firstName()) // Tom
	// fmt.Printf("%s\n", ai.Name()) // ai.Name undefined (type Namer has no field or method Name) 结构体自己的方法不能通过接口调用
	ai = &Chinese{}
	fmt.Printf("%s\n", ai.firstName()) // 曹
}

// 有的时候，也会以一种稍微不同的方式来使用接口这个词：从某个类型的角度来看，它的接口指的是：它的所有导出方法，只不过没有显式地为这些导出方法额外定一个接口而已。
