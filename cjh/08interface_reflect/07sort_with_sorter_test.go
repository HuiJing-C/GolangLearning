package _8interface_reflect

// 对于基本类型的排序，标准库已经提供了相关的排序函数，所以不需要我们再重复造轮子了。对于一般性的排序，sort 包定义了一个接口：
//
// type Interface interface {
//    Len() int
//    Less(i, j int) bool
//    Swap(i, j int)
// }
// 这个接口总结了需要用于排序的抽象方法，函数 Sort(data Interface) 用来对此类对象进行排序，可以用它们来实现对其他类型的数据（非基本类型）进行排序。
