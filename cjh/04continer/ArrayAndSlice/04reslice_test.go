package ArrayAndSlice

import (
	"fmt"
	"testing"
)

func TestReSliceOne(t *testing.T) {
	// 在切片到达容量上限(相关数组长度)之前，我们可以对其进行阔扩容，注意，只能扩容
	// 改变切片长度的过程称之为 切片重组 reslicing
	slice1 := make([]int, 5, 10)
	fmt.Printf("%d %d\n", len(slice1), cap(slice1)) // 5 10
	slice1 = slice1[0 : len(slice1)+1]
	fmt.Printf("%d\n", len(slice1)) // 6
	slice1 = slice1[0:cap(slice1)]
	fmt.Printf("%d\n", len(slice1)) // 10
	// slice1 = slice1[0: cap(slice1) + 1] // panic: runtime error: slice bounds out of range [:11] with capacity 10

	fmt.Printf("%d\n", len(slice1[2:2])) // 0
	fmt.Printf("%d\n", len(slice1[2:3])) // 1
}
