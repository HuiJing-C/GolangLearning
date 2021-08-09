package ArrayAndSlice

import (
	"fmt"
	"testing"
)

func TestCopyOne(t *testing.T) {
	// 如果想要增加切片的容量，我们必须建一个更大切片，然后把原来切片的数据全部复制过去，用到copy(dst, src []type)函数，src覆盖dst
	slice_from := []int{1, 2, 3}
	slice_to := make([]int, 10)
	copy(slice_to, slice_from)
	fmt.Printf("%v %p %d\n", slice_to, &slice_to, cap(slice_to)) // [1 2 3 0 0 0 0 0 0 0] 0xc00008e4e0 10
	slice_to = append(slice_to, 10, 11)
	fmt.Printf("%v %p %d\n", slice_to, &slice_to, cap(slice_to)) // [1 2 3 0 0 0 0 0 0 0 10 11] 0xc0000f0000 20
	slice_to = append(slice_to, []int{12, 13}...)
	fmt.Printf("%v %p %d\n", slice_to, &slice_to, cap(slice_to)) // [1 2 3 0 0 0 0 0 0 0 10 11 12 13] 0xc0000f0000 20
}

// func append(s[]T, x ...T) []T 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；
// 追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。
// 因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。
// 如果你想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可：x = append(x, y...)
