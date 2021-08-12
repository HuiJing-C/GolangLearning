package _8interface_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 在 04struct_with_tag_test.go 节我们看到可以通过反射来分析一个结构体。
// 进一步探讨强大的反射功能。反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，
// 例如它的大小、方法和 动态 的调用这些方法。这对于没有源代码的包尤其有用。这是一个强大的工具，除非真得有必要，否则应当避免使用或小心使用

// 变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。
func TestReflect(t *testing.T) {
	var x float64 = 3.14
	fmt.Printf("%v\n", reflect.TypeOf(x))  // float64
	fmt.Printf("%v\n", reflect.ValueOf(x)) // 3.14

	typeOf := reflect.TypeOf(x)
	println(typeOf.Kind().String()) // float64
	valueOf := reflect.ValueOf(x)
	println(valueOf.Kind() == reflect.Float64) // true
	// 变量 v 的 Interface() 方法可以得到还原（接口）值，所以可以这样打印 v 的值
	fmt.Println(valueOf.Interface()) // 3.14

	// 假设我们要把 x 的值改为 3.1415。Value 有一些方法可以完成这个任务，但是必须小心使用：v.SetFloat(3.1415)
	// valueOf.SetFloat(3.1415) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	// ====================不好理解，后续探究========================
}
