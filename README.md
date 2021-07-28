# GolangLearning
go语言自学



纠错笔记：

1：想使用单测的方式来进行语言学习，但是发现以下形式并不能在IDEA运行

```
func TestA(t *testing.T) {
}
```

Answer：Go语言单测只能是以test结尾的才可以使用



2：同一project下不同package之间不可以互相import

Answer：由于上一步全是以test结尾的.go文件，Go语言不导出package下以test结尾的公开方法或者属性
