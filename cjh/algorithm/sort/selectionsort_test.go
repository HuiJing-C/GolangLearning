package sort

import (
	"fmt"
	"testing"
)

// 选择排序
// 选一个最小的放第一位，从剩余的中再选一个最小的放第二位...
func TestSelectionSort(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 9, 8, 7, 6, 0}
	for i := 0; i < len(s)-1; i++ {
		hasSorted := true
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
				hasSorted = false
			}
		}
		fmt.Printf("%v\n", s)
		if hasSorted {
			break
		}
	}
}
