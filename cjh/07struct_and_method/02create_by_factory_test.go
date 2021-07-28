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

// 如果想强制使用工厂方法，使结构提变成私有的。可以运用可见性规则
type peop struct {
	name string
	age  int
}

func NewPeop(name string, age int) *peop {
	return &peop{"cjh", 18}
}
