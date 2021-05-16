package control_struct

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// go中的switch语句接受任何形式的表达式

func TestSwitchOne(t *testing.T) {
	// switch var1 {
	// case val1:
	// 	...
	// case val2,val3:// 匹配多个条件
	// 	...
	// default:
	// 	...
	// }
	// val1,val2,val3可以使同类型的任意值，前{ 必须和switch在同一行

	// 程序自上而下匹配，一旦匹配到，后续的就不会再执行，而且不需要特别使用break表示结束
	// 如果想执行后续分支，可以使用fallthrough关键字

	// rand.Intn(10)产生的是伪随机数，不管调多少次，值都是一样的，因为他的默认资源就是单一的
	// 所以必须调用rand.Seed()传入一个变化的值，就可以生成时刻变化的值
	rand.Seed(time.Now().UnixNano())
	var i int = rand.Intn(10)
	println(i)
	switch i {
	case 0, 1, 2:
	default: // default放在这里和放在最后效果是一样的，只有非0，1，2，3,4,5时才会执行这里
		time.Sleep(time.Second)
	case 3, 4, 5:
		os.Exit(1)
	}

	// case语句之后，不需要使用花括号将多行语句括起来，但可以在分支中进行任意形式的编码。当代码块只有一行时，可以直接放置在 case 语句之后。

	// 还可以在case语句中添加return语句来提前结束代码块，并且还要在switch后添加相应return语句确保函数始终有返回值

	// default分支可以出现在任何顺序中，但是最好放在最后，因为他的作用类似于if-else中的else,表示不符合任何已给条件
}

func TestSwitchTwo(t *testing.T) {
	// 第二种用法是不提供任何判断值，实际默认为true，然后在每个case中测试不同的条件、
	// 当任一分支测试结果为true时，该分支代码就会被执行
	// switch {
	// case condition1:
	// 	...
	// case condition2:
	// 	...
	// default:
	// 	...
	// }
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	println(i)
	switch {
	case i < 10:
		fmt.Printf("A %d\n", i)
		fallthrough // 会不加判定的执行下一个case语句
	case i < 6:
		fmt.Printf("B %d\n", i)
		fallthrough
	case i < 4:
		fmt.Printf("C %d\n", i)
	default:
		println("666")
	}
}

func TestSwitchThree(t *testing.T) {
	// switch第三种形式是带一个初始化的表达式，可以非常优雅的进行条件判断，并且可以进行多值赋值
	// switch initialization; {//注意这个分号
	// case val1:
	// 	...
	// case val2:
	// 	...
	// default:
	// 	...
	// }
	rand.Seed(time.Now().UnixNano())
	switch a, b := rand.Intn(10), rand.Intn(10); {
	case a < b:
		println("a<b")
	case a > b:
		println("a>b")
	default:
		println("a==b")
	}
}
