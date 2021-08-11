package _map

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	for k, v := range m {
		fmt.Printf("%s === %d\n", k, v)
		v = v * 2
	}
	// one === 1
	// two === 2
	fmt.Printf("%v\n", m) // map[one:1 two:2]  for range 内部是实际值的拷贝，改变其值并不能改变map
}

// 如果想创建一个map类型的切片，要使用两次make，第一次make切片，第二次make map
func TestMake(t *testing.T) {
	slice1 := make([]map[string]string, 5)
	for i := range slice1 {
		slice1[i] = make(map[string]string)
		slice1[i]["one"] = "1"
	}
	fmt.Printf("%v\n", slice1) // [map[one:1] map[one:1] map[one:1] map[one:1] map[one:1]]

	slice2 := make([]map[string]string, 5)
	for _, v := range slice2 {
		v = make(map[string]string) // 不能这样赋值，因为v只是实际值的拷贝
		v["one"] = "1"
	}
	fmt.Printf("%v\n", slice2) // [map[] map[] map[] map[] map[]]
}

// map内部既不是按key排序的也不是按value排序的，如果想对map排序，可以先将key赋给一个切片，把切片排序，使用切片的值来取map的值
