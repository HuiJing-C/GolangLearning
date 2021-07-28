package error

import (
	"fmt"
	"testing"
)

// 不能随意地用 panic 中止程序，必须尽力补救错误让程序能继续执行

// 在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。
// 这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 panicking。
func TestRPOne(t *testing.T) {
	fmt.Printf("%s\n", "start")
	panic("666")
	fmt.Printf("%s\n", "end")
}
