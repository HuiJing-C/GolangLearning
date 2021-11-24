package sort

import (
	"fmt"
	"testing"
)

// 冒泡
// 依次比较相邻两元素，小的放前面。
func TestBubble(t *testing.T) {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < len(s)-1; i++ {
		noneSort := true
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				noneSort = false
			}
		}
		fmt.Printf("%v\n", s)
		if noneSort {
			break
		}
	}
}
