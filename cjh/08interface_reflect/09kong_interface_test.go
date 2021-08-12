package _8interface_reflect

// 空接口或者最小接口 不包含任何方法，它对实现不做任何要求
// 任何其他类型都实现了空接口（它不仅仅像 Java/C# 中 Object 引用类型），any 或 Any 是空接口一个很好的别名或缩写
// 每个 interface {} 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针

// >>>>>>>>>>>>>复制数据切片至空接口切片
// 假设你有一个 myType 类型的数据切片，你想将切片中的数据复制到一个空接口切片中，类似：
//
// var dataSlice []myType = FuncReturnSlice()
// var interfaceSlice []interface{} = dataSlice
// 可惜不能这么做，编译时会出错：cannot use dataSlice (type []myType) as type []interface { } in assignment。
// 原因是它们俩在内存中的布局是不一样的（参考 Go wiki）。
// 必须使用 for-range 语句来一个一个显式地赋值：
// var dataSlice []myType = FuncReturnSlice()
// var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
// for i, d := range dataSlice {
//    interfaceSlice[i] = d
// }
