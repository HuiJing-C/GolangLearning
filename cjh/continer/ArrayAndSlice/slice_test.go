package ArrayAndSlice

import (
	"fmt"
	"testing"
)

// 切片是对数组一个连续片段的引用（该数组称之为相关数组），因此切片是一个引用类型
// 切片是可索引的，并且可以使用len()获取长度
// 给定项的索引可能比相关数组的相同元素的索引小
// 和数组不同，切片的长度是可以在运行时修改的，最小为0最大为相关数组的长度：--切片是一个长度可变的数组
// cap()可以测量切片的最长长度；他等于切片长度 + 数组除切片之外的长度 cap(s)就是s[0]到数组末尾的数组长度，所以下式恒成立 0 <= len(s) <= cap(s)

func TestSliceOne(t *testing.T) {
	ints := [5]int{1, 2, 3, 4, 5}
	s1 := ints[0:3]
	fmt.Printf("%v %d %d\n", s1, len(s1), cap(s1)) // [1 2 3] 3 5
	// 不同切片如果表示同一个数组的片段，他们可以共享数据。不同数组总是代表不同的存储。数组实际上是切片的构建块
	s2 := ints[1:4]
	s2[0] = 3
	fmt.Printf("%v %d %d\n", s1, len(s1), cap(s1)) // [1 3 3] 3 5

	// 切片的优点：因为是引用，所以不需要使用额外的内存，并且比数组更有效率
	// 声明格式：var s []type (不需要说明长度)，一个切片在未初始化之前默认值是 nil,长度为0
	var s3 []int
	fmt.Printf("%t %d\n", s3 == nil, len(s3)) // true 0
	// 初始化格式：var s []type = arr[start:end] 这表示切片是从start - end-1索引元素构成的子集
	// 全集表达式 var s []type = arr[:] / arr[0:len(arr)] / &arr
	var s4 []int = ints[0:len(ints)]
	fmt.Printf("%v\n", s4) // [1 3 3 4 5]
	// 去除切片最后一个元素 s[:len(s)-1]

	// 一个由123组成的切片可以这么生成
	// s5 := [3]int{1, 2, 3}
	// s6 := []int{1, 2, 3}

	// s1 := s[:]是用切片组成切片，拥有相同的元素，但仍指向相同的相关数组
	s7 := s4[:]
	s7[0] = 7
	fmt.Printf("%v\n", s4) // [7 3 3 4 5]

	// 一个数组可以扩展到他的大小上限 s = s[:cap(s)]，如果再扩大就要导致运行时错误
	// s7 = s7[:cap(s7)+1] // runtime error: slice bounds out of range [:6] with capacity 5

	// 对于所欲切片，以下状态总是成立的。i是一个整数且: 0 <= i <= len(s)
	// s == s[:i] + s[i:]    len(s) <= cap(s)

	// 切片也可以用类似数组的初始化方式 var s = []int{1,2,3,4,5}，这样就创建了一个长度为5的数组并且创建了一个相关切片
	var s8 = []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", s8) // [1 2 3 4 5]
}

func TestSliceTwo(t *testing.T) {
	// 切片在内存中实际是一个有3个域的结构体：指向相关数组的指针，切片长度以及切片容量
	// s是一个slice,后移一位操作：s = s[1:]，切片只能后移，前移s[-1:]会导致编译报错，切片不能重新分片以获取数组的前一个元素
	// 注意：千万不要用指针指向切片，切片本身就是一个引用类型了，切片本身就是一个指针

	// 当一个函数要对数组进行操作时，可能一直需要将入参转换为切片，当调用该函数时，将数组分片，创建一个切片引用并传递给该函数
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", Sum(arr[:])) // 15
}
func Sum(s []int) (i int) {
	for _, v := range s {
		i += v
	}
	return
}

func TestSliceThree(t *testing.T) {
	// 当没有相关数组是我们可以使用make([]type, len)来构建一个切片，并构建相关数组。这里的len既是数组的长度也是slice的初始长度
	var ints []int = make([]int, 5)                          // or ints := make([]int, 5) 此时cap(ints) == len(ints) == 5
	fmt.Printf("%t %d\n", len(ints) == cap(ints), len(ints)) // true 5
	// 如果不想切片占用整个数组，可以使用make([]type, len, cap)。其中cap是可选参数
	ints2 := make([]int, 5, 10)                   // 此方法与new([100]int)[0:50]可以生成相同的切片
	fmt.Printf("%d %d\n", len(ints2), cap(ints2)) // 5 10
}

// new()和make()的区别，二者都在堆上分配内存，但是行为不同，适用于不同的类型
// new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体,它相当于 &T{}
// make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
