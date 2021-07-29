package goroutine_channel

import (
	"fmt"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

// 如果我们不在 main() 中等待，协程会随着程序的结束而消亡。
// 当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束。
// 这就是为什么在服务器程序中，每一个请求都会启动一个协程来处理，server() 函数必须保持运行状态。通常使用一个无限循环来达到这样的目的。
// 协程是独立的处理单元，一旦陆续启动一些协程，你无法确定他们是什么时候真正开始执行的。你的代码逻辑必须独立于协程调用的顺序。

/*
协程更有用的一个例子应该是在一个非常长的数组中查找一个元素。
将数组分割为若干个不重复的切片，然后给每一个切片启动一个协程进行查找计算。这样许多并行的协程可以用来进行查找任务，整体的查找时间会缩短（除以协程的数量）
*/

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

// Go 协程（goroutines）和协程（coroutines）
// （“Go协程（goroutines）” 即是指的是 Go 语言中的协程。而“协程（coroutines）”指的是其他语言中的协程概念，仅在本目录中出现。）
//
// 在其他语言中，比如 C#，Lua 或者 Python 都有协程的概念。这个名字表明它和 Go协程有些相似，不过有两点不同：
//
// Go 协程意味着并行（或者可以以并行的方式部署），协程一般来说不是这样的
// Go 协程通过通道来通信；协程通过让出和恢复操作来通信
// Go 协程比协程更强大，也很容易从协程的逻辑复用到 Go 协程。
