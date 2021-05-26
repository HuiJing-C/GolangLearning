package ArrayAndSlice

import (
	"fmt"
	"testing"
)

func TestForRangeOne(t *testing.T) {
	ints := [5]int{1, 2, 3, 4, 5}
	slice1 := ints[:]
	// for range 结构可以用于数组和切片
	// index是切片的索引，value是切片索引位置的值的拷贝，不能用来修改切片的值
	for index, value := range slice1 {
		fmt.Printf("%d %d\n", index, value)
		value = value * 2
	}
	fmt.Printf("%v\n", slice1) // [1 2 3 4 5]
}
