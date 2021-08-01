package struct_and_method

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// 在 Go 语言中，结构体就像是类的一种简化形式，那么面向对象程序员可能会问：类的方法在哪里呢？
// 在 Go 中有一个概念，它和方法有着同样的名字，并且大体上意思相同：Go 方法是作用在[接收者(receiver)]上的一个函数，接收者是某种类型的变量。
// 因此方法是一种特殊类型的函数。
//
// 接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
// 但是接收者不能是一个接口类型（参考 第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；
// 如果这样做会引发一个编译错误：invalid receiver type…。
// 最后接收者不能是一个指针类型(*指针)，但是它可以是任何其他允许类型的指针(*其他类型)。

// 一个类型加上它的方法等价于面向对象中的一个类。
// 一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件
// 唯一的要求是：它们必须是同一个包的。
//
// 类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的 方法集。
//
// 因为方法是函数，所以同样的，>>不允许方法重载<<，即对于一个类型只能有一个给定名称的方法。
// 但是如果基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的：
// func (a *denseMatrix) Add(b Matrix) Matrix
// func (a *sparseMatrix) Add(b Matrix) Matrix

// >>> 别名类型不能有它原始类型上已经定义过的方法 <<<

// 一般定义如下
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

// 如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。
// 如果 recv 一个指针，Go 会自动解引用。
// 如果方法不需要使用 recv 的值，可以用 _ 替换它
// recv 就像是面向对象语言中的 this 或 self，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 this 或self 作为 receiver 的名字。

type TwoInts struct {
	a, b int
}

func (this *TwoInts) Add() int {
	return this.a + this.b
}

func TestMethod(t *testing.T) {
	tt := &TwoInts{1, 2}
	fmt.Printf("%d\n", tt.Add()) // 3

	fmt.Printf("%d\n", IntVector{1, 2, 3}.Sum()) // 6
}

// 下面是非结构体类型上方法的例子
type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, vl := range v {
		s += vl
	}
	return
}

// 下面这段代码有什么错？
/*
package main

import "container/list"

func (p *list.List) Iter() {
    // ...
}

func main() {
    lst := new(list.List)
    for _= range list.Iter() {
    }
}
*/

// 类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float 或类似这些的类型上定义方法。
// 试图在 int 类型上定义方法会得到一个编译错误：cannot define new methods on non-local type int

// 但是有一个绕点的方式：可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。
// 或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func TestMethod2(t *testing.T) {
	m := myTime{time.Now()}
	fmt.Printf("%s\n", m.String())      // 2021-08-01 17:07:57.903227 +0800 CST m=+0.000562140
	fmt.Printf("%s\n", m.first3Chars()) // 202
}

// 函数与方法的区别
// 函数将变量作为参数：Function1(recv)
// 方法在变量上被调用：recv.Method1()
// 在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。

// 在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。
// 方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

// 疑难点：指针 或者 值 作为接受者
/*
鉴于性能的原因，recv 最常见的是一个指向 receiver_type 的指针（因为我们不想要一个实例的拷贝，如果按值调用的话就会是这样），特别是在 receiver 类型是结构体时，就更应如此。
如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。
看下例
注意 Go 为我们做了探测工作，我们自己并没有指出是是否在指针上调用方法，Go 替我们做了这些事情。b1 是值而 b2 是指针，方法都支持运行了。
*/

type BC struct {
	thing int
}

func (b *BC) change() { b.thing = 1 }

func (b BC) write() string { return fmt.Sprint(b) }

func (b BC) write2() string {
	b.thing = 2
	return fmt.Sprint(b)
}

func TestMethod3(t *testing.T) {
	var b1 BC // b1是值
	b1.change()
	fmt.Println(b1.write()) // {1}

	b2 := new(BC) // b2是指针
	b2.change()
	fmt.Println(b2.write()) // {1}
}

// 试着在 write2() 中改变接收者b的值：将会看到它可以正常编译，但是开始的 b 没有被改变。

func TestMethod4(t *testing.T) {
	b := BC{3}
	fmt.Printf("%v\n", b.write2()) // {2}
	fmt.Printf("%v\n", b)          // {3}
}

// 方法并不是一定需要指针作为接收者，但是如果这么做代价有点昂贵
// 因为接受者作为值传递给方法，那么传递的是它的拷贝，这在 Go 中合法的。
// 如果是在这个类型的指针上调用次方法，会自动解引 -- >>>指针方法和值方法都可以在指针或非指针上被调用<<<

// 方法和未导出字段
// 对于未导出的字段，采用实例.字段名的方式是错误的。如何解决这个问题？
// 可以考虑面向对象语言的get/set方法，对于set要有Set前缀，对于get，只要成员名就可以了
type Person struct {
	name string
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func TestMethod5(t *testing.T) {
	p := new(Person)
	p.SetName("cjh")
	fmt.Printf("%s\n", p.Name()) // cjh
}

// 并发访问对象
// 对象的字段（属性）不应该由 2 个或 2 个以上的不同线程在同一时间去改变。
// 如果在程序发生这种情况，为了安全并发访问，可以使用包 sync（参考06package）中的方法。在第 11goroutine_channel 中会通过 goroutines 和 channels 探索另一种方式。

// 内嵌类型的方法和继承
// 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 "继承" 了这些方法：将父类型放在子类型中来实现亚型。
// 这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，也类似 Ruby 中的混入（mixin）。
type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) Work() {
	c.Start()
	c.Stop()
}

// 内嵌结构体上的方法可以直接在外层类型的实例上调用
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) Abs2() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

func TestMethod6(t *testing.T) {
	n := &NamedPoint{Point{3, 4}, "cjh"}
	fmt.Println(n.Abs()) // 5
}

// 内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法“晋升”成为了外层类型的方法。
// 当然类型可以有只作用于本身实例而不作用于内嵌“父”类型上的方法。可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
func (n *NamedPoint) Abs2() float64 {
	return n.Point.Abs() * 100.
}

func TestMethod7(t *testing.T) {
	n := &NamedPoint{Point{3, 4}, "cjh"}
	fmt.Println(n.Abs2()) // 500
}

// 因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承，就像：type Child struct { Father; Mother}
// 结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法。

// >>>如何在类型中嵌入功能<<<
// A：聚合（或组合）：包含一个所需功能类型的具名字段 - A中包含一个B类型的字段
// B：内嵌：内嵌（匿名地）所需功能类型 - A中包含一个B的匿名字段 >>> 如果内嵌类型嵌入了其他类型，也是可以的，那些类型的方法可以直接在外层类型中使用。

type Volume struct {
	NamedPoint
	high float64
}

func (v *Volume) Vol() float64 {
	return v.x * v.y * v.high
}

func TestMethod8(t *testing.T) {
	volume := Volume{
		NamedPoint: NamedPoint{Point{1, 2}, "cjh"},
		high:       3,
	}
	fmt.Printf("%f %f %f\n", volume.Vol(), volume.Abs(), volume.Abs2()) // 6.000000 2.236068 223.606798
}

// >>> 多重继承 <<<
// 多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++ 和 Python 例外）。
// 因为在类继承层次中，多重继承会给编译器引入额外的复杂度。但是在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。如下
/*type A struct {
	B
	C
}*/

// >>> 通用方法和方法命名 <<<
// 在编程中一些基本操作会一遍又一遍的出现，比如打开（Open）、关闭（Close）、读（Read）、写（Write）、排序（Sort）等等
// 并且它们都有一个大致的意思：打开（Open）可以作用于一个文件、一个网络连接、一个数据库连接等等。具体的实现可能千差万别，但是基本的概念是一致的。
// 在 Go 语言中，通过使用接口（参考 第 11 章），标准库广泛的应用了这些规则，在标准库中这些通用方法都有一致的名字，比如 Open()、Read()、Write()等。
// 想写规范的 Go 程序，就应该遵守这些约定，给方法合适的名字和签名，就像那些通用方法那样。这样做会使 Go 开发的软件更加具有一致性和可读性。
// 比如：如果需要一个 convert-to-string 方法，应该命名为 String()，而不是 ToString()

// Go 不需要一个显式的类定义，如同 Java、C++、C# 等那样，相反地，“类”是通过提供一组作用于一个共同类型的方法集来隐式定义的。类型可以是结构体或者任何用户自定义类型。
// 比如：我们想定义自己的 Integer 类型，并添加一些类似转换成字符串的方法，在 Go 中可以如下定义：
//
// type Integer int
// func (i *Integer) String() string {
//    return strconv.Itoa(i)
// }
// 在 Java 或 C# 中，这个方法需要和类 Integer 的定义放在一起，在 Ruby 中可以直接在基本类型 int 上定义这个方法。

/*
总结
在 Go 中，类型就是类（数据和关联的方法）。Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。
在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程。
许多开发者说相比于类继承，Go 的接口提供了更强大、却更简单的多态行为。

备注
如果真的需要更多面向对象的能力，看一下 goop 包（Go Object-Oriented Programming），它由 Scott Pakin 编写: 它给 Go 提供了 JavaScript 风格的对象（基于原型的对象），
并且支持多重继承和类型独立分派，通过它可以实现你喜欢的其他编程语言里的一些结构。
*/
