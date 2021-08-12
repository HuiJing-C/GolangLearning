package _8interface_reflect

import (
	"fmt"
	"testing"
)

// 作用于变量上的方法实际上是不区分变量到底是指针还是值的。
// 当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的，幸运的是，如果使用不当编译器会给出错误。
type List []int

func (l List) Len() int {
	return len(l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func TestMethods(t *testing.T) {
	// A bare value
	var lst List
	// CountInto(lst, 1, 10)
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}
	//  在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个 Appender，而它的方法 Append 只定义在指针上。
	//  在 lst 上调用 LongEnough 是可以的，因为 Len 定义在值上。
	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
	// 	在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender，并且它的方法 Append 定义在指针上。
	// 	在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
}

// 总结

// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
// ---指针方法可以通过指针调用
// ---值方法可以通过值调用
// ---接收者是值的方法可以通过指针调用，因为指针会首先被解引用
// ---接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
// ---将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

// Go 语言规范定义了接口方法集的调用规则：
//
// 类型 T 的可调用方法集包含接受者为 *T 或 T 的所有方法集（因为指针会被自动解指引）
// 类型 T 的可调用方法集包含接受者为 T 的所有方法
// 类型 *T 的可调用方法集不包含接受者为 T 的方法 （因为指针会被自动解指引，所以看起来好像*T可以调用接受者为值的方法集）看03type_assertion_test.go - 27

// 指针可以调指针与值，值不可以调指针
