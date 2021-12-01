package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertion(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Printf("开始%v\n", s)
	for i := 1; i < len(s); i++ {
		current := s[i]
		for j := 0; j <= i; j++ {
			if current <= s[j] {
				moveAfterIndex(s[:i+1], j)
				s[j] = current
				break
			}
		}
		fmt.Printf("第%d次排序结果%v\n", i, s)
	}
	fmt.Printf("结果%v\n", s)
}

func moveAfterIndex(s []int, index int) {
	l := len(s)
	if l < 2 || index > l-1 {
		return
	}
	// for i := l; i > index && i > 1; i-- {
	// 	s[i-1] = s[i-2]
	// }
	s = append(s[:index+1], s[index:len(s)-1]...)
}
