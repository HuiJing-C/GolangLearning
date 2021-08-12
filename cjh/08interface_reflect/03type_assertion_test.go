package _8interface_reflect

import "testing"

// 一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型。
// 在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：
type E interface {
	read()
}
type F struct{}

func (f *F) read() {
	println("F read")
}

func (f *F) write() {
	println("F write")
}

func TestTypeAssertion(t *testing.T) {
	var e E  // e必须是一个接口类型的变量，不然会报错：invalid type assertion: e.(*F) (non-interface type (type of e) on left)
	e = &F{} // 如果少了&就会报错：cannot use F{} (type F) as type E in assignment: F does not implement E (read method has pointer receiver)
	f := e.(*F)
	f.read()
	f.write()
	// 如果忽略 e.(*F) 中的 * 号，会导致编译错误：impossible type assertion: F does not implement E (read method has pointer receiver)，看下例
	e = G{} // 如果加&就会报错：interface conversion: _8interface_reflect.E is *_8interface_reflect.G, not _8interface_reflect.G [recovered]
	g := e.(G)
	g.read() // G read
	// ??? 为什么，为什么一个接口类型的变量，如果结构体用的是指针方法就必须用指针赋值，结构体是实例方法，就必须用实例赋值

	// 类型断言可能是无效的，虽然编译器会尽力检查转换是否有效，但是它不可能预见所有的可能性。
	// 如果转换在程序运行时失败会导致错误发生。更安全的方式是使用以下形式来进行类型断言：
	if f2, ok := e.(*F); ok {
		f2.read()
	}
	// 如果转换合法，f2 是 e 转换到类型 F 的值，ok 会是 true；否则 e 是类型 F 的零值，ok 是 false，也没有运行时错误发生。
	// 多数情况下，我们可能只是想在 if 中测试一下 ok 的值，此时使用以下的方法会是最方便的：
	//
	// if _, ok := e.(*F); ok {
	//    // ...
	// }
}

type G struct{}

func (g G) read() {
	println("G read")
}
