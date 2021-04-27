package main

import (
	"fmt"
	"os"
)

func main() {
	student := &Student{"zhangsan", 18}
	student.name = "lisi"
	fmt.Printf("PID : %d %d %v", os.Getpid(), Add(1, 2), student)
}
