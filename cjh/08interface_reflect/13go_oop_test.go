package _8interface_reflect

// Go是如何实现传统OOP语言的 封装，继承，多态
// 封装
// 可见性规则
// 1）包范围内的：通过标识符首字母小写，对象 只在它所在的包内可见
//
// 2）可导出的：通过标识符首字母大写，对象 对所在包以外也可见
// 类型只拥有自己所在包中定义的方法

// 继承
// 通过内嵌多个想要行为的类型实现

// 多态
// 用接口实现，某个类型的实例可以赋值给他所实现的任意接口类型的变量
// 类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。Go 接口不是 Java 和 C# 接口的变体，而且接口间是不相关的，并且是大规模编程和可适应的演进型设计的关键。