package pointer

import (
	"fmt"
	"testing"
)

// 取地址符：&
func TestOne(t *testing.T) {
	var i1 = 5
	fmt.Printf("i1 value is: %d, location in memory: %p\n", i1, &i1) // i1 value is: 5, location in memory: 0xc00001a260(变化的十六进制)

	// 指针定义
	var intP *int
	intP = &i1 // intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1。
	fmt.Printf("%p\n", intP)

	// 指针在不同机器上占用的字节不一样，32位机器占4字节，64位机器占8字节，并且与他指向的值的大小无关
	// 可以在指针前加*来取指针所指向的值，*号是一个类型变更器
	// 当一个指针没有分配到任何变量时，他的值是0x0
	// var p *type p后一定要加空格，因为var p*type是合法的，在复杂表达式中容易被误认为是一个乘法表达式
	var intA *int
	fmt.Printf("%p\n", intA) // 0x0
	intA = &i1
	fmt.Printf("%d\n", *intA) // 5
	// 对任一个变量var,以下表达式恒等 var == *(&var)
}

func TestTwo(t *testing.T) {
	s := "Hello World!"
	var p *string = &s
	*p = "cjh"
	fmt.Printf("pointer p: %p\n", p)  // 0xc00004c510
	fmt.Printf("string *p: %s\n", *p) // cjh
	fmt.Printf("string s: %s\n", s)   // cjh
	// 通过对 *p 赋另一个值来更改“对象”，这样 s 也会随之更改。

	// 不能到的一个文字或者常量的地址
	// const con = 5
	// ptr := &con  cannot take the address of con
	// ptr2 := &10  cannot take the address of 10

	// 指针的高级应用是可以传递一个变量的引用，这样不会传递变量的拷贝，很廉价，减少内存占用和提高效率

	// 对于一个空指针的反向引用是不合法的，会导致程序崩溃
	var ptr3 *int = nil
	*ptr3 = 0 // runtime 10error: invalid memory address or nil pointer dereference [recovered]
}
