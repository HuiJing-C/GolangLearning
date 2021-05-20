package function

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestClosePkg(t *testing.T) {
	// 当我们不想给函数起名的识货，就可以使用匿名函数(闭包：函数式编程术语)
	f := func(x, y int) int { return x + y } // 这样的函数不能独立存在，不然会报：non-declaration statement outside function body
	i := f(1, 2)
	fmt.Printf("%d\n", i) // 3

	// 也可以直接对匿名函数进行调用,直接跟()
	i2 := func(a int) int { return a * a }(2)
	fmt.Printf("%d\n", i2) // 4

	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %p\n", g, g)
	}
	// 0  - g is of type func(int) and has value 0x50c250
	// 1  - g is of type func(int) and has value 0x50c250
	// 2  - g is of type func(int) and has value 0x50c250
	// 3  - g is of type func(int) and has value 0x50c250
	// 可以看到函数g代表的是func(int),变量的值是一个内存地址
}

func TestClosePkg2(t *testing.T) {
	// 关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。
	fmt.Println(f()) // 2
	// 这可用于在返回语句之后修改返回的 error 时使用

	// 匿名函数还可以配合 go 关键字来作为 goroutine 使用
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

// 闭包的应用
// 它们被允许调用定义在其它环境下的变量。闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态
// 一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁
// 闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装
// 另一个不错的应用就是使用闭包来完成更加简洁的错误检查

// 下面是两个返回func(int) int的函数
func Add2() func(a int) int {
	return func(a int) int {
		return a + 2
	}
}

func Addr(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func TestClosePkg3(t *testing.T) {
	// 函数Addr2()被赋值到变量f中（类型为func(int)int）
	f := Addr2()
	fmt.Printf("%d\n", f(1))
	fmt.Printf("%d\n", f(10))
	fmt.Printf("%d\n", f(100))
	// 1 11 111
	// 在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 10 = 11，最后 11 + 100 = 111
	// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	// 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的

	var g int
	go func(i int) {
		s := 0
		for j := 1; j <= i; j++ {
			s += j
		}
		g = s
	}(100) // Passes argument 1000 to the function literal.
	time.Sleep(time.Second)
	fmt.Printf("%d\n", g) // 5050

	jpg := SufferFactory(".jpg")
	gif := SufferFactory(".gif")
	fmt.Printf("%s\n", jpg("test"))     // test.jpg
	fmt.Printf("%s\n", gif("test"))     // test.gif
	fmt.Printf("%s\n", gif("test.gif")) // test.gif
}

func Addr2() func(b int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// 返回一组相似函数的函数称为函数工厂或者高阶函数，下面是给文件添加后缀的函数
func SufferFactory(suffer string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffer) {
			return name + suffer
		}
		return name
	}
}
