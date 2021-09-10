package main

import "github.com/HuiJing-C/GolangLearning/cjh/private_utils"

func main() {
	println(private_utils.SliceToString([]interface{}{1, 2, 3}, ", "))
	println(private_utils.SliceToString2([]interface{}{1, 2, 3}, ", "))
}
