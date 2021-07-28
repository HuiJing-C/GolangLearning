package error

import (
	"fmt"
	"testing"
)

func TestRecoverOne(t *testing.T) {
	// （recover）内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。
	// recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值
	// 如果是正常执行，调用 recover 会返回 nil，且没有其它效果
	fmt.Printf("%s\n", "test start")
	analogPanic()
	fmt.Printf("%s\n", "test end")
	// test start
	// panic recover 666
	// test end
}

// Go 库的原则是即使在包的内部使用了 panic，在它的对外接口（API）中也必须用 recover 处理成返回显式的错误。

func analogPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover %v\n", err)
		}
	}()
	badCall()
	fmt.Printf("%s\n", "after panic")
}

func badCall() {
	panic("666")
}
