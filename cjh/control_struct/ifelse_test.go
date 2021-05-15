package control_struct

import (
	"fmt"
	"runtime"
	"testing"
)

func TestOne(t *testing.T) {
	// 通用模式
	condition := false
	condition2 := true
	if condition {
		// do something
	} else if condition2 {
		// do something
	} else {
		// catch all or default
	}
	// else结构不要太多，即使没办法，也要尽可能把先满足条件的放在前面

	// 注意：即使代码只有一行，大括号也不可以省略

	// 条件表达式括号可以省略，但是当条件比较复杂时，可以用()让代码更易读，或者提高某个运算的优先级
	// 当if结构中有break, continue, goto, return时，常见用法是省略else部分
	// if condition {
	// 	return 0
	// }
	// return 1
}

func TestTwo(t *testing.T) {
	// 一些有用的例子
	// 1:判断字符串为空
	// if str == "" {...}

	// 判断操作系统，这段代码一般被放在 init() 函数中执行
	if runtime.GOOS == "windows" {

	} else { // Unix like
	}

	// if语句可以有一个初始化语句(给一个变量赋值)，具有以下固定格式
	// if initialization; condition {
	// }
	if val := 10; val > 6 {
		// ...
	}
	// 声明的变量val，作用域在if结构中，如果val在if结构前就已经存在，该变量原来的值就会被隐藏
	fmt.Printf("")
}
