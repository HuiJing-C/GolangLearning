package function

import (
	"fmt"
	"testing"
)

// defer允许我们将执行推迟到函数返回之前(任意return之后，为什么是之后？因为return并不是只返回值，还会做一些计算)
// defer有点类似于面向对象语言java c#中的finally,常用来释放某些资源
// 多个defer被注册时，会逆序执行，类似栈的FILO
func TestDeferOne(t *testing.T) {
	fmt.Printf("%s\n", "In")
	defer Defer1("defer1")
	defer Defer1("defer2")
	defer Defer1("defer3")
	fmt.Printf("%s\n", "Out")
	// In
	// Out
	// defer3
	// defer2
	// defer1
}

// defer允许我们进行一些资源的收尾工作，比如：
// 1:关闭文件流  open a file defer file.Close()
// 2:解锁一个加锁的资源 mu.Lock() defer mu.Unlock()
// 3:打印最终报告 printHeader() defer printFooter()
// 4:关闭数据库连接 open a database connection defer disconnectFromDB()
func Defer1(str string) {
	fmt.Printf("%s\n", str)
}

// 一个实用功能是：实用defer来追踪进入函数和离开函数，形式如下
// func a() {
// 	A()
// 	defer b()
// }

// 另一种调试的手法是使用defer来打印函数的入参与返回值，如下：
// func func1(s string) (n int, err error) {
//    defer func() {
//        log.Printf("func1(%q) = %d, %v", s, n, err)
//    }()
//    return 7, io.EOF
// }

// 下面是一些内置函数，他们不需要导入就可以直接使用
// close:用于管道通信

// len、cap:len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）

// new、make:new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用户内置引用类型（切片、map 和管道）。
// 它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。
// 它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作。new() 是一个函数，不要忘记它的括号

// copy、append:用于复制和连接切片

// panic、recover:两者均用于错误处理机制

// print、println:底层打印函数，在部署环境中建议使用 fmt 包

// complex、real、imag:用于创建和操作复数
