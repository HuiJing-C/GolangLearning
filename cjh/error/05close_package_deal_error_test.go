package error

import (
	"fmt"
	"testing"
)

/*
每当函数返回时，应该检查是否有错误发生：但是这会导致重复乏味的代码。结合 defer/panic/recover 机制和闭包可以得到一个更加优雅的模式。
不过这个模式只有当所有的函数都是同一种签名时可用，这样就有相当大的限制。一个很好的使用它的例子是 web 应用，所有的处理函数都是下面这样：
func handler1(w http.ResponseWriter, r *http.Request) { ... }

假设所有的函数都有这样的签名：
func f(a type1, b type2)
参数的数量和类型是不相关的。

我们给这个类型一个名字：

fType1 = func f(a type1, b type2)
在模式中使用了两个帮助函数：
1）check：这是用来检查是否有错误和 panic 发生的函数：
func check(err error) { if err != nil { panic(err) } }

2）errorhandler：这是一个包装函数。接收一个 fType1 类型的函数 fn 并返回一个调用 fn 的函数。里面就包含有 defer/recover 机制
func errorHandler(fn fType1) fType1 {
    return func(a type1, b type2) {
        defer func() {
            if e, ok := recover().(error); ok {
                log.Printf(“run time panic: %v”, err)
            }
        }()
        fn(a, b)
    }
}

当错误发生时会 recover 并打印在日志中；除了简单的打印，应用也可以用 template 包为用户生成自定义的输出。check() 函数会在所有的被调函数中调用，像这样：

func f1(a type1, b type2) {
    ...
    f, _, err := // call function/method
    check(err)
    t, err := // call function/method
    check(err)
    _, err2 := // call function/method
    check(err2)
    ...
}
通过这种机制，所有的错误都会被 recover，并且调用函数后的错误检查代码也被简化为调用 check(err) 即可。
在这种模式下，不同的错误处理必须对应不同的函数类型；它们（错误处理）可能被隐藏在错误处理包内部。可选的更加通用的方式是用一个空接口类型的切片作为参数和返回值。
*/

func TestClosePackageError(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
