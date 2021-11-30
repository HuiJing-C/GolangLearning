package sort

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	s := []int{9, 7, 8, 6, 5, 4, 3, 2, 1}
	fmt.Printf("===%v\n", s)
	for i := 1; i < len(s); i++ {
		for j := range s[:i] {
			if s[i] < s[j] {
				s = append(append(s[:j], s[i]), s[i+1:]...)
			}
		}
		fmt.Printf("%v\n", s)
	}
	fmt.Printf("===%v\n", s)
}
