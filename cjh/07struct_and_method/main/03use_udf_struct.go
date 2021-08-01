package main

import (
	"fmt"
	s "github.com/HuiJing-C/GolangLearning/cjh/07struct_and_method"
)

func main() {
	b := new(s.BBB)
	b.Name = "cjh"
	b.Age = 18
	fmt.Printf("%s\n", b.Name)
	fmt.Printf("%d\n", b.Age)
}
