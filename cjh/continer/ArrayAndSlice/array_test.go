package ArrayAndSlice

import (
	"fmt"
	"testing"
)

func TestArrayOne(t *testing.T) {
	// 数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列（这是一种同构的数据结构）
	// 数组长度必须是一个常量表达式，并且必须是一个非负整数。数组长度也是数组类型的一部分，所以[5]int和[10]int是属于不同类型的。数组的编译时值初始化是按照数组顺序完成的
	// 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型。当使用值时我们必须先做一个类型判断
	// 数组长度最大为 2Gb

	// 声明格式：var var_name [len]type,
	// 数组索引从0开始，可以通过索引来访问其中的元素与修改，所以数组是可变的
	// 数组在初始化时会给每个元素赋零值
	var ints [5]int
	fmt.Printf("%v\n", ints) // [0 0 0 0 0]
	// 索引错误：runtime error: index out of range
	for index, _ := range ints {
		ints[index] = index * index
	}
	fmt.Printf("%v\n", ints)      // [0 1 4 9 16]
	fmt.Printf("%d\n", len(ints)) // 5
}

func TestArrayTwo(t *testing.T) {
	// 还可以通过以下方式直接初始化数组 arr := [...]{val1, val2, val3 ...}
	ints := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v %d\n", ints, len(ints)) // [1 2 3] 3
	// go语言中数组是值类型，所以可以通过new()来创建，但是此时new出来的是指向数组的指针
	var ints1 = new([5]int)
	fmt.Printf("%v\n", ints1) // &[0 0 0 0 0]
	// 此时ints1类型为*[5]int，ints类型是[5]int
	// ints1 = ints // cannot use ints (type [5]int) as type *[5]int in assignment
	ints2 := ints // 此时赋值，在内存中进行了一次值拷贝，所以ints2此时是ints拷贝后的值，修改ints2并不会修改ints中的值
	ints2[0] = 2
	fmt.Printf("%v\n", ints)  // [1 2 3 4 5]
	fmt.Printf("%v\n", ints2) // [2 2 3 4 5]
	// 如果想修改赋值后的ints2可以使用取地址符&来取得ints的地址，以引用方式传递过来
}

func TestArrayThree(t *testing.T) {
	// 如果数组值已经提前知道了，那么可以通过 数组常量 的方法来初始化数组，有以下几种方式
	var arr1 = [5]int{1, 2, 3}
	fmt.Printf("%v %d\n", arr1, len(arr1)) // [1 2 3 0 0] 5
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v %d\n", arr2, len(arr2)) // [1 2 3 4 5] 5
	var arr3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("%v %d\n", arr3, len(arr3)) // [1 2 3 4 5] 5
	var arr4 = [5]int{1: 6, 2: 7}
	fmt.Printf("%v %d\n", arr4, len(arr4)) // [0 6 7 0 0] 5
	var arr5 = []int{1: 6, 2: 7}
	fmt.Printf("%v %d\n", arr5, len(arr5)) // [0 6 7] 3
}

func TestArrayFour(t *testing.T) {
	// 多维数组
	var screen [3][3]int
	fmt.Printf("%v\n", screen) // [[0 0 0] [0 0 0] [0 0 0]]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			screen[i][j] = i * j
		}
	}
	fmt.Printf("%v\n", screen) // [[0 0 0] [0 1 2] [0 2 4]]
}

// 将数组传递给函数会消耗大量内存，所以有以下两种方式优化1：使用数组的指针 2：使用切片。1不常见，多使用2见slice_test.go
func TestArrayFive(t *testing.T) {
	arr := &[5]int{1, 2, 3}
	fmt.Printf("%d\n", SumFun(arr)) // 6
}
func SumFun(s *[5]int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
