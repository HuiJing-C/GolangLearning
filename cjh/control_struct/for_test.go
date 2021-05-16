package control_struct

import (
	"fmt"
	"strings"
	"testing"
)

func TestForOne(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	// ASCII编码的字符占一个字节，可以用索引遍历字符。但是非ASCII编码的字符（2-4字节）不能单纯的使用索引来判断是否为同一个字符
	str := "ab"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d\t%c###", i, str[i])
	}
	// 0	a###1	b

	str2 := "中国"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d\t%c###", i, str2[i])
	}
	// 0	ä###1	¸###2	­###3	å###4	###5	½

	println()
	// 打印倒三角
	for i, j := 10, 0; i > 0; i, j = i-2, j+1 {
		fmt.Printf("%s%s%s\n", strings.Repeat(" ", j), strings.Repeat("*", i), strings.Repeat(" ", j))
	}
}

func TestForTwo(t *testing.T) {
	// 基于条件判断的迭代
	// for condition {}
	i := 5
	for i > 0 {
		i--
		fmt.Printf("%d\n", i)
	}
}

func TestForThree(t *testing.T) {
	// 当for后没有判定条件时，会认为条件永远为：true  可以是如下形式
	// for i:=0;;i++ {}
	// for ;; {}
	// for {}
	// 死循环循环体内必须有判定条件以确保循环体会在某一时刻退出，使用break(退出循环体)、return(提前对函数体进行返回)退出循环
}

func TestForRang(t *testing.T) {
	// for-rang结构可以迭代任何一个集合（包括数组与map）
	// 一般形式如下：for index, value := range coll { }
	// 注意：value始终为对应索引的 "值拷贝"，因此他只有只读性质。对它做任何修改都不会影响集合中的原有值
	// 如果value为指针，则会产生指针拷贝，依旧可以修改集合中的原值
	strs := []string{"A", "b"}
	fmt.Printf("%v\n", strs) // [A b]
	for _, value := range strs {
		value = value + ":"
		fmt.Printf("%s\n", value) // A: B:
	}
	fmt.Printf("%v\n", strs) // [A b]

	peoples := []*People{{name: "zhangsan", age: 12}, {name: "lisi", age: 13}}
	fmt.Printf("%s\t%s\n", peoples[0].name, peoples[1].name) // zhangsan	lisi
	for _, people := range peoples {
		people.name = people.name + "2"
	}
	fmt.Printf("%s\t%s\n", peoples[0].name, peoples[1].name) // zhangsan2	lisi2

	// 每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符
	// 一个字符串是 Unicode 编码的字符（或称之为 rune）集合
	str := "ab中国"
	for _, value := range str {
		fmt.Printf("%c\n", value) // a b 中 国
	}
}

type People struct {
	name string
	age  uint8
}
