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

type student struct {
	Name string
	Age  int
}

// map rang的一个坑
func TestMapRange(t *testing.T) {
	// 定义map
	m := make(map[string]*student)

	// 定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// 将数组依次添加
	for _, stu := range stus { // stu是同一个地址里的值，只不过在每次循环时被赋予了不同的值，所以每次&stu都是指向同一个地址，所以遍历完后m的所有value都是wang
		fmt.Printf("%p %s\n", &stu, stu.Name) // %p打印后为同一个地址
		m[stu.Name] = &stu                    // 指针拷贝
	}

	// 打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
	// zhou => wang
	// li => wang
	// wang => wang

	println("========================以上是错误的示例=======================")

	// >>>>>>>>>>>> 可以这样改
	m2 := make(map[string]student)
	for _, stu := range stus { // 虽然stu还是同一个地址的不同值，但是下一行是进行的值拷贝，所以并没有将地址信息带过去
		m2[stu.Name] = stu
	}

	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}
	// zhou => zhou
	// li => li
	// wang => wang

	println("======================第二种方法====================")
	// >>>>>>>>> 也可以这样改
	m3 := make(map[string]student)
	for i := 0; i < len(stus); i++ {
		m3[stus[i].Name] = stus[i]
	}

	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}
}
