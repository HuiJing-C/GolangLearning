package control_struct

import (
	"fmt"
	"testing"
)

func TestTags(t *testing.T) {
	// 标签：某一行第一个以冒号：结尾的单词，gofmt会将后续代码自动移动至下一行；
	// 标签大小写敏感，为了提升可读性，一般建议全部使用大写字母
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 4 {
				continue LABEL1 // j==4,5都不会打印出来
			}
			fmt.Printf("i is %d j is %d\n", i, j)
		}
	}
	// 当标签配合break使用时，会退出离标签最近的循环体
	// 标签还可以配合goto语句模拟循环

	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
	// 01234

	// 使用标签和goto语句是不被鼓励的，他们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求

	// 如果必须使用goto,应当使用正序的标签(标签位于goto语句之后)，但注意标签和goto语句之间不能出现定义新变量的语句，否则会导致编译报错
}
