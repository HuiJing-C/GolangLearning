package struct_and_method

import (
	"fmt"
	"testing"
)

// 类型 Stu 在定义它的包 struct_and_method 中必须是唯一的，它的完全类型名是：struct_and_method.Stu

// 声明
// type structName struct {
// 	field1 type1
// 	field2 type2
// }
// type structName struct {field1, field2 type}

// 对于从来不用的field名称可以用_代替

// 可以声明一个结构体类型变量，然后像下面这种赋值
func TestOne(t *testing.T) {
	var s Stu
	s.name = "cjh"
	s.age = 18
	fmt.Printf("%v\n", s)

	var s2 Stu     // s2是Stu
	s3 := new(Stu) // s3是*Stu
	// fmt.Printf("%p\n", s2) // Printf format %p has arg s2 of wrong type command-line-arguments.Stu
	fmt.Println(s3)                            // &{ 0}
	fmt.Printf("%+v ||| %+v %p\n", s2, s3, s3) // s2, s3都是其结构体的零值 {name: age:0} ||| &{name: age:0} 0xc00000c060
}

type Stu struct {
	name string
	age  int
}

// 数组也可以看做一种特殊的结构体，只不过访问属性是用下标而不是属性名

// go中可以使用(结构体).(字段名)来引用结构体的字段，这种做法叫做选择器（selector）,无论是结构体类型还是结构体的指针类型，都可以使用
// new(Type)等价于&Type{}
func TestTwo(t *testing.T) {
	s := &Stu{
		name: "cjh",
		age:  18,
	}
	// 此时s类型是*Stu
	// 或者更简单的做法
	s = &Stu{"cjh1", 19}
	fmt.Printf("%v\n", s) // &{cjh1 19}

	// 或者指定字段名赋值,可以指定某些字段而不是全部
	s = &Stu{age: 20, name: "cjh2"}
	fmt.Printf("%v\n", s) // z&{cjh2 20}

	// 可以通过解指针的方式来访问结构体属性
	(*s).name = "cjh3"
	fmt.Printf("%v\n", s) // &{cjh3 20}
}

// Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

// 结构体有别名alias的话，底层都是一种类型，可以直接显式转化
type SS Stu

func TestThree(t *testing.T) {
	s := SS{"cjh", 18}
	s2 := Stu(s)
	fmt.Printf("%v\n", s2) // {cjh 18}
}
