package goroutine_channel

import (
	"fmt"
	"testing"
)

// 通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。
// 只有发送者需要关闭通道，接收者永远不会需要。

func TestCloseChannel(t *testing.T) {
	// 如何关闭管道？
	// 可以通过函数 close(ch) 来完成：这个将通道标记为无法通过发送操作 <- 接受更多的值；
	// 给已经关闭的通道发送或者再次关闭都会导致运行时的 panic。在创建一个通道后使用 defer 语句是个不错的办法
	ch := make(chan int)
	defer close(ch)

	// 如何来检测可以收到没有被阻塞（或者通道没有被关闭）
	// 可以使用逗号，ok 操作符：用来检测通道是否被关闭。
	// v, ok := <-ch   // ok is true if v received value

	ch2 := make(chan string)
	go sendData2(ch2)
	getData2(ch2)

	// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：

	ch3 := make(chan string)
	go sendData2(ch3)
	for s := range ch3 {
		println(s)
	}
}

func sendData2(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"
	ch <- "E"
	close(ch)
}

func getData2(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s\n", input)
	}

}
