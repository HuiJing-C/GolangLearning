package goroutine_channel

import (
	"fmt"
	"testing"
	"time"
)

// time 包中有一些有趣的功能可以和通道组合使用。
// 其中就包含了 time.Ticker 结构体，这个对象以指定的时间间隔重复的向通道 C 发送时间值：
// 在协程周期性的执行一些事情（打印状态日志，输出，计算等等）的时候非常有用。
// 调用 Stop() 使计时器停止，在 defer 语句中使用。这些都很好的适应 select 语句:

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1e9)
	defer ticker.Stop()

	for i := 0; i < 10; i++ {
		select {
		case <-ticker.C:
			fmt.Printf("666\n")
		default:
			fmt.Printf("777\n")
		}
		time.Sleep(1e9)
	}
}

// time.Tick() 函数声明为 Tick(d Duration) <-chan Time，当你想返回一个通道而不必关闭它的时候这个函数非常有用：
// 它以 d 为周期给返回的通道发送时间，d是纳秒数。如果需要像下边的代码一样，限制处理频率
func TestTick(t *testing.T) {
	var dur time.Duration = 1e9
	chRate := time.Tick(dur) // a tick every 1 second
	for {
		<-chRate // rate limit
		println(666)
	}
	// 这样只会按照指定频率处理请求：chRate 阻塞了更高的频率。每秒处理的频率可以根据机器负载（和/或）资源的情况而增加或减少
}

// 定时器（Timer）结构体看上去和计时器（Ticker）结构体的确很像（构造为 NewTimer(d Duration)），但是它只发送一次时间，在 Dration d 之后。

// 还有 time.After(d), 在 Duration d 之后，当前时间被发到返回的通道；所以它和 NewTimer(d).C 是等价的；它类似 Tick()，
// 但是 After() 只发送一次时间。下边有个很具体的示例，很好的阐明了 select 中 default 的作用
func TestTickAndTimer(t *testing.T) {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

// 习惯用法：简单超时模式
// 要从通道 ticker 中接收数据，但是最多等待1秒。先创建一个信号通道，然后启动一个 lambda 协程，协程在给通道发送数据之前是休眠的
// 然后使用 select 语句接收 ticker 或者 ch 的数据：如果 ticker 在 1 秒内没有收到数据，就选择到了 ch 分支并放弃了 ticker 的读取
func TestSimpleTimeOut(t *testing.T) {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		ch <- true
	}()
	ticker := time.NewTicker(2e9)
	select {
	case <-ticker.C:
		fmt.Printf("A\n")
	case <-ch:
		fmt.Printf("timeout\n")
		break
	}
}
