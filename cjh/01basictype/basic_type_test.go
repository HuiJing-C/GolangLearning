package basictype

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// go语言没有float类型，只有float32/float64
// float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。由于精确度的缘故，使用 == 或者 != 来比较浮点数时应当非常小心。
// 最好在正式使用前测试对于精确度要求较高的运算。
// 你应该尽可能地使用 float64，因为 math 包中所有有关数学运算的函数都会要求接收这个类型。
// 可以通过前缀加0表示8进制 012 = 10
// 0x表示十六进制 0x12 = 18
// e表示10的连乘 1e3 = 1000,6.002e12 = 6.002 * 1e12
func TestOne(t *testing.T) {
	i := uint64(12) // 指定类型并赋值
	//fmt.Println(i)
	println(i)

	// Go 中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少，因此允许常量之间的混合使用
	var a int
	var b int32
	a = 15
	//b = a + a    // 编译错误
	b = b + 5 // 因为 5 是常量，所以可以通过编译
	fmt.Println(a, b)

	// go无法隐式转换，都要显式转换
	var f1 float32 = 12
	var f2 float64
	//f2 = f1 //cannot use f1 (type float32) as type float64 in assignment
	// 要显式转换
	f2 = float64(f1)
	fmt.Println(f1, f2)
}

func TestTwo(t *testing.T) {
	// %d用于格式化整数 %x/%X用于格式化16进制表示的数字
	a := uint64(12)
	b := 0xFF
	fmt.Printf("%d %x\n", a, b) //12 ff

	// %g格式化浮点数 %f输出浮点数 %e科学计数法输出浮点数
	c := float64(3.1415926)
	fmt.Printf("%g %f %e\n", c, c, c) //3.1415926 3.141593 3.141593e+00

	// %n.m输出n并保留m位小数，除了使用 g 之外，还可以使用 e 或者 f
	fmt.Printf("%5.2e\n", c)
	fmt.Printf("%5.2g\n", c)
	fmt.Printf("%5.2f\n", c)
	// %v格式化复数complex64/complex128
	re := float32(1.2)
	im := float32(3.3)
	comp := complex(re, im)
	comp2 := 5 + 3i
	fmt.Printf("%v\n", comp)
	fmt.Printf("%v\n", comp2)

	// %0d输出定长的整数
	i := int64(1234567890)
	fmt.Printf("%0d", i)

	/*
	   %v 输出结构体 {10 30}
	   %+v 输出结构体显示字段名 {one:10 tow:30}
	   %#v 输出结构体源代码片段 main.Point{one:10, tow:30}
	   %T 输出值的类型            main.Point
	   %t 输出格式化布尔值      true
	   %d`输出标准的十进制格式化 100
	   %b`输出标准的二进制格式化 99 对应 1100011
	   %c`输出定整数的对应字符  99 对应 c
	   %x`输出十六进制编码  99 对应 63
	   %f`输出十进制格式化  99 对应 63
	   %e`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	   %E`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	   %s 进行基本的字符串输出   "\"string\""  对应 "string"
	   %q 源代码中那样带有双引号的输出   "\"string\""  对应 "\"string\""
	   %p 输出一个指针的值   &jgt 对应 0xc00004a090
	   % 后面使用数字来控制输出宽度 默认结果使用右对齐并且通过空格来填充空白部分
	   %2.2f  指定浮点型的输出宽度 1.2 对应  1.20
	   %*2.2f  指定浮点型的输出宽度对齐，使用 `-` 标志 1.2 对应  *1.20
	*/
}

func TestThree(t *testing.T) {
	//随机数
	for i := 0; i < 10; i++ {
		in := rand.Int()
		fmt.Printf("%d / ", in)
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		in := rand.Intn(10)
		fmt.Printf("%d / ", in)
	}
	fmt.Println()

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}
