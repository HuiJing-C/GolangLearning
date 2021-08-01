package struct_and_method

import (
	"fmt"
	"testing"
)

// go没有常见的面向对象的语言的那种构造函数，但是可以很容易实现"构造函数"的理念，一般会定义一个New开头的工厂方法
type Peo struct {
	name string
	age  int
}

func NewPeo(name string, age int) *Peo {
	return &Peo{
		name: name,
		age:  age,
	}
}

func TestFactory(t *testing.T) {
	peo := NewPeo("cjh", 12)
	fmt.Printf("%v\n", peo) // &{cjh 12}
}

// 如果想强制使用工厂方法，使结构体变成私有的。可以运用可见性规则
// 通过应用可见性规则，就可以禁止使用 new() 函数，强制用户使用工厂方法，从而使类型变成私有的，就像在面向对象语言中那样。
type peop struct {
	name string
	age  int
}

func NewPeop(name string, age int) *peop {
	return &peop{"cjh", 18}
}

// make只适用于三种类型，切片slice，字典map，管道channel
// 与new()的区别可以从下面的案例区分
type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func TestNewAndMake(t *testing.T) {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// NOT OK
	// z := make(Bar) // 编译错误：cannot make type Bar
	// (*y).thingOne = "hello"
	// (*y).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	u := new(Foo)         // 返回的一个指向nil的指针，他尚未被分配内存
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}

// 试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！
// 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
