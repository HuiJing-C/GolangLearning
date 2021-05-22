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

func TestAliceTwo(t *testing.T) {
	// 切片在内存中实际是一个有3个域的结构体：指向相关数组的指针，切片长度以及切片容量
}
