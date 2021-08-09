package _map

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	m := make(map[string]int)
	m["one"] = 1
	// 判断键是否存在，不存在的话值为零值
	v, contain := m["one"]
	if contain {
		println(v) // 1
	}
	v2, contain2 := m["two"]
	if !contain2 {
		println(v2) // 0
	}
}

func TestDelete(t *testing.T) {
	m := map[string]int{"one": 1}
	delete(m, "one")
	fmt.Printf("%v\n", m) // 存在，删除
	delete(m, "two")
	fmt.Printf("%v\n", m) // 不存在，也不会产生错误
}
