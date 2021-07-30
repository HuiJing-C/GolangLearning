package goroutine_channel

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 协程可以使用共享变量来通信，但是很不提倡这样做，因为这种方式给所有的共享内存的多线程都带来了困难。
// Go有一个特殊的类型，通道（channel），像是通道（管道），可以通过它们发送类型化的数据在协程之间通信，可以避开所有内存共享导致的坑；
// 通道的通信方式保证了同步性。数据通过通道：同一时间只有一个协程可以访问数据：所以不会出现数据竞争，设计如此。数据的归属（可以读写数据的能力）被传递。

// 通常使用这样的格式来声明通道：var identifier chan datatype
// 未初始化的通道的值是nil。
// 所以通道只能传输一种类型的数据，比如 chan int 或者 chan string，所有的类型都可以用于通道，空接口 interface{} 也可以。甚至可以（有时非常有用）创建通道的通道。

// 通道实际上是类型化消息的队列：使数据得以传输。它是先进先出（FIFO）结构的，所以可以保证发送给他们的元素的顺序
// （有些人知道，通道可以比作 Unix shells 中的双向管道（tw-way pipe））。
// 通道也是引用类型，所以我们使用 make() 函数来给它分配内存。这里先声明了一个字符串通道 ch1，然后创建了它（实例化）：
func TestChanOne(t *testing.T) {
	/*var ch1 chan string
	ch1 = make(chan string)*/
	// 或者
	ch2 := make(chan string)
	// 甚至是
	ch3 := make(chan func())
	fmt.Printf("%v | %v\n", ch2, ch3)
}

// 所以通道是对象的第一类型：可以存储在变量中，作为函数的参数传递，从函数返回以及通过通道发送它们自身。另外它们是类型化的，允许类型检查，比如尝试使用整数通道发送一个指针。

// 通信操作符 ->/<-信息按照箭头的方向流动
// 流向通道（发送）
// ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

// 从通道流出（接收），三种方式：
// int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。
// <- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：
/*
if <- ch != 1000{
    ...
}
*/
// 操作符 <- 也被用来发送和接收，Go 尽管不必要，为了可读性，通道的命名通常以 ch 开头或者包含 chan。通道的发送和接收操作都是自动的：它们通常一气呵成。下面的示例展示了通信操作。
func TestChanTwo(t *testing.T) {
	ch := make(chan string)
	// sendData(ch)
	// getData(ch) // fatal error: all goroutines are asleep - deadlock!

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9) // ABC
	// 如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
	// 注释掉sleep, sendData就没机会输出
	/*
		协程之间的同步非常重要：
		main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
		getData() 使用了无限循环：它随着 sendData() 的发送完成和 ch 变空也结束了。
		如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic：fatal error: all goroutines are asleep - deadlock!

		为什么会这样？运行时会检查所有的协程（也许只有一个是这种情况）是否在等待（可以读取或者写入某个通道），意味着程序无法处理。
		这是死锁（deadlock）形式，运行时可以检测到这种情况。
		注意：不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。

		如果删掉sleep, 在函数 getData() 的一开始插入 time.Sleep(1e9)，不会出现错误但也没有输出: 因为协程会随着程序的结束而消亡
	*/
}

func sendData(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
}

func getData(ch chan string) {
	var intput string
	for {
		intput = <-ch
		print(intput)
	}
}

// 通道阻塞
/*默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。可以想象一个无缓冲的通道在没有空间来保存数据的时候：
必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的：
1）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：
新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。
2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
尽管这看上去是非常严格的约束，实际在大部分情况下工作的很不错。*/

// 下面程序验证了以上理论，一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字 0。
func TestBlockChannel(t *testing.T) {
	ch := make(chan int)
	go pump(ch) // 此行注释掉就会出现 fatal error: all goroutines are asleep - deadlock! ---原因：主协程消费管道数据被阻塞, TestA可以避免
	fmt.Printf("%d\n", <-ch)
}

func TestA(t *testing.T) {
	c := make(chan int)
	go func() { c <- 1 }()
	go func() { fmt.Printf("%d\n", <-c) }() // 此种情况两个协程调换顺序都可以成功，因为主协程没被阻塞
	time.Sleep(1e8)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 为通道解除阻塞定义了 suck 函数来在无限循环中读取通道
func suck(ch chan int) {
	for {
		fmt.Printf("%d\n", <-ch)
	}
}

func TestBlockChannel2(t *testing.T) {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	time.Sleep(1e8)
	fmt.Printf("===%d\n", <-ch) // ===15722
}

// 通信是一种同步形式：通过通道，两个协程在通信（协程会和）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
// 甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go 运行时会检查并 panic，停止程序。死锁几乎完全是由糟糕的设计导致的。如下例
// 无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。
func TestBlockChannel3(t *testing.T) {
	ch := make(chan int)
	ch <- 1   // 先生产，生产阻塞，没有生产。再消费，因为没有生产出来，所以消费也阻塞
	go f1(ch) // 和上一行换一下就可以运行了
}

func f1(ch chan int) {
	fmt.Printf("%d\n", <-ch)
}

/*同步通信---带缓冲的通道*/

func TestBufferChannel(t *testing.T) {
	// 在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。
	// 缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的cap 函数可以返回缓冲区的容量。
	ch := make(chan int, 10)
	fmt.Printf("%d\n", cap(ch)) // 10
	for i := 0; i < 10; i++ {
		go func() { ch <- i }()
		// time.Sleep(5e7) // 不加sleep，管道中就可能会被写入重复的值，聚体原因参考TestParallel()
		// ch <- i // 或者这样写就对了
	}
	// ch <- 10 // fatal error: all goroutines are asleep - deadlock!
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", <-ch)
	}
}

/*如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功。
ch :=make(chan type, value)
value == 0 -> synchronous, unbuffered (阻塞）同步
value > 0 -> asynchronous, buffered（非阻塞）异步 取决于value元素
若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable）。
要在首要位置使用无缓冲通道来设计算法，只在不确定的情况下使用缓冲。*/

/*在其他协程运行时让 main 程序无限阻塞的通常做法是在 main 函数的最后放置一个{}。
也可以使用通道让 main 程序等待协程完成，就是所谓的信号量模式*/
func TestBlockMain(t *testing.T) {
	ch := make(chan int)
	/*信号量模式
	协程通过在通道 ch 中放置一个值来处理结束的信号。main 协程等待 <-ch 直到从中获取到值。*/
	go func() {
		time.Sleep(1e9)
		ch <- 1
	}()
	<-ch

	// 或者等待两个协程完成，每一个都会对切片s的一部分进行排序，片段如下：
	done := make(chan bool)
	// doSort is a lambda function, so a closure which knows the channel done:
	doSort := func() {
		// sort(s)
		println("sort")
		done <- true
	}
	go doSort()
	go doSort()
	<-done
	<-done
}

func TestParallel(t *testing.T) {
	/*用信号量控制多个协程并行计算*/
	a := make([]int, 10)
	b := make([]int, 10)
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(in int) {
			b[in] = doSum(a)
			ch <- true
		}(i)
	}
	/*注意闭合：i 是作为参数传入闭合函数的，从外层循环中隐藏了变量 i 。让每个协程有一份 i 的拷贝；如果不是作为参数传入闭合函数，会出现各种奇怪的问题
	另外，for 循环的下一次迭代会更新所有协程中 i 的值。切片 res 没有传入闭合函数，因为协程不需要单独拷贝一份。
	切片 b 也在闭合函数中但并不是参数。*/
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Printf("%v\n", b) // [47 46 39 34 44 44 43 43 31 43]
}

func doSum(s []int) (sum int) {
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		s[i] = rand.Intn(10)
		sum += s[i]
	}
	fmt.Printf("%v\n", s)
	return
}

// producer and consumer
func TestPC(t *testing.T) {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)
	<-done
}

// integer producer:
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

// channel range
func TestChannelRange(t *testing.T) {
	/*从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。*/
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for i := range ch {
		fmt.Printf("%d\n", i)
	}
}

// 习惯用法:通道迭代模式
// 通常，需要从包含了地址索引字段 items 的容器给通道填入元素。为容器的类型定义一个方法 Iter()，返回一个只读的通道
type container struct {
	Items []int
}

func (c *container) Len() int {
	return len(c.Items)
}

func (c *container) Iter() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < c.Len(); i++ {
			ch <- c.Items[i]
		}
		close(ch) // 写完要关，不然读的时候会死锁
	}()
	return ch
}

// 在协程里，一个 for 循环迭代容器 c 中的元素（对于树或图的算法，这种简单的 for 循环可以替换为深度优先搜索）。
// 调用这个方法的代码可以这样迭代容器：
func TestItems(t *testing.T) {
	c := &container{Items: []int{1, 2, 3}}
	for i := range c.Iter() {
		fmt.Printf("%d\n", i)
	}
}

// 习惯用法：生产者消费者
// 假设你有 Produce() 函数来产生 Consume 函数需要的值。它们都可以运行在独立的协程中，生产者在通道中放入给消费者读取的值。整个处理过程可以替换为无限循环
/*
for {
    Consume(Produce())
}
*/

// 通道的方向:
// 通道类型可以用注解来表示它只发送或者只接收：
// var send_only chan<- int        // channel can only send data 只能向通道发数据（只写）
// var recv_only <-chan int        // channel can onley receive data 只能从通道读取数据 （只读）
// 只接收的通道（只读）（<-chan T）无法关闭，因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。
// 通道创建的时候都是双向的，但也可以分配有方向的通道变量，就像以下代码：
func TestDirect(t *testing.T) {
	var c = make(chan int) // bidirectional
	go source(c)
	go sink(c)
}

func source(ch chan<- int) {
	for {
		ch <- 1
	}
}

func sink(ch <-chan int) {
	for {
		<-ch
	}
}

// 习惯用法：管道和选择器
// 更具体的例子还有协程处理它从通道接收的数据并发送给输出通道：
func TestChannelSelect(t *testing.T) {
	sendChan := make(chan int)
	receiveChan := make(chan string)
	go processChannel(sendChan, receiveChan)
}
func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		result := strconv.FormatInt(int64(inValue), 0)
		out <- result
	}
}

// 通过使用方向注解来限制协程对通道的操作。
// 这里有一个来自 Go 指导的很赞的例子，打印了输出的素数，使用选择器（‘筛’）作为它的算法。每个 prime 都有一个选择器
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ { // Error:这里设置i < 100就会报fatal error: all goroutines are asleep - deadlock!原因暂时不知道，后续补充
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
// filter(in, out chan int, prime int) 拷贝整数到输出通道，丢弃掉可以被 prime 整除的数字。然后每个 prime 又开启了一个新的协程，生成器和选择器并发请求
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

func TestSuShu(t *testing.T) {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// 下面这个版本引入了上边的习惯用法：函数 sieve、generate 和 filter 都是工厂；它们创建通道并返回，而且使用了协程的 lambda 函数。
// TestSuShu2 函数现在短小清晰：它调用 sieve() 返回了包含素数的通道，然后通过 fmt.Println(<-primes) 打印出来。

func TestSuShu2(t *testing.T) {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}

// Send the sequence 2, 3, 4, ... to returned channel
func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}
