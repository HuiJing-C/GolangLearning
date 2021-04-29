package main

/*
author:cjh
*/
// my first golang project
import (
	"fmt"
	"os"
)

const PI = 3.14

//别名
type INT int64

type Student struct {
	Name string
	Age  uint8
}

func AddAndDelete(a, b int8) (c, d int8) {
	c = a + b
	d = a - b
	return
}

func main() {
	student := &Student{"zhangsan", 18}
	student.Name = "lisi"
	add, _ := AddAndDelete(1, 2)
	format := fmt.Sprintf("PID : %d %d %v %f", os.Getpid(), add, student, PI)

	fmt.Println(format)
	var tmpInt INT = 10
	//显式转换
	i := int(tmpInt)
	fmt.Println(tmpInt)
	fmt.Println(i)
}
