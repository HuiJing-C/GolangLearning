package _map

import (
	"fmt"
	"testing"
)

func TestMapOne(t *testing.T) {
	// 声明格式如下，在声明的时候不需要知道 05map 的长度，05map 是可以动态增长的，它不存在固定长度或者最大限制。
	// 如果想指定长度，可以使用make(05map[type]type, len)
	// 当 05map 增长到容量上限的时候，如果再增加新的 key-value 对，05map 的大小会自动加 1。
	// 所以出于性能的考虑，对于大的 05map 或者会快速扩张的 05map，即使只是大概知道容量，也最好先标明。
	// var map1 05map[key_type]value_type
	// 未初始化的 05map 的值是 nil。
	var map1 map[string]string
	fmt.Printf("%t %d\n", map1 == nil, len(map1)) // true 0

	// key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key，但是指针和接口类型可以。
	// 如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
	// value 可以是任意类型的；通过使用空接口类型，可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言

	// 05map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。
	// 通过 key 在 05map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。
	// 05map 也可以用函数作为自己的值，这样就可以用来做分支结构：key 用来选择要执行的函数。

	// 取值 v := 05map[key] 将key对应的值赋值给v，如果没有对应的值，v为value对应类型的零值
	// 赋值 05map[key] = value

	//  05map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。
	var map2 map[string]int
	map2 = map[string]int{"one": 1, "two": 2}
	map3 := make(map[string]float32)
	map3["PI"] = 3.1415926
	fmt.Printf("%v %v\n", map2, map3) // 05map[one:1 two:2] 05map[PI:3.1415925]

	// 不要使用 new，永远用 make 来构造 05map
	// 注意 如果错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	// m := new(05map[string]int)
	// m["a"] = 1 // invalid operation: m["a"] (type *05map[string]int does not support indexing)
}

// 既然一个 key 只能对应一个 value，而 value 又是一个原始类型，那么如果一个 key 要对应多个值怎么办？
// 例如，当我们要处理unix机器上的所有进程，以父进程（pid 为整形）作为 key，所有的子进程（以所有子进程的 pid 组成的切片）作为 value。
// 通过将 value 定义为 []int 类型或者其他类型的切片，就可以优雅的解决这个问题。
// mp1 := make(05map[int][]int)
// mp2 := make(05map[int]*[]int)
