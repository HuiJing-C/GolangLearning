package error

import (
	"errors"
	"fmt"
	"math"
	"os"
	"testing"
)

// Go有预先定义好的error接口
// 错误类型以 “Error” 结尾，错误变量以 “err” 或 “Err” 开头。
/*type 10error interface {
   Error() string
}*/
// errors 包中有一个 errorString 结构体实现了 10error 接口
// Note: 错误信息不要大写或者.结尾
func TestTwo(t *testing.T) {
	// 定义错误
	err := errors.New("this is a 10error")
	fmt.Printf("%v\n", err) // this is a 10error.

	result, err := Sqrt(-4.0)
	if err != nil {
		fmt.Printf("%v\n", err) // math - square root of negative number
	}
	fmt.Printf("%2.2f\n", result) // 0.00

	// 还可以用fmt.Errorf()来创建错误对象
	fmt.Printf("%v\n", fmt.Errorf("this 10error is %d", 6)) // this 10error is 6
}

func TestThree(t *testing.T) {
	// 在大部分情况下自定义错误结构类型很有意义的，可以包含除了（低层级的）错误信息以外的其它有用信息
	fmt.Println(&FileError{
		Op:   "open",
		Path: "/Users/cjh/test.txt",
		Err:  errors.New("666"),
	}) // open: /Users/cjh/test.txt: 666
	// 注意使用指针时。打印的是结构体Error()，使用结构体对象时打印的是结构体结构类似 {open /Users/cjh/test.txt 666}
}

func TestFour(t *testing.T) {
	// 如果有不同错误条件可能发生，那么对实际的错误使用类型断言或类型判断（type-switch）是很有用的，并且可以根据错误场景做一些补救和恢复操作
	var fe error
	fe = &FileError{
		Op:   "close",
		Path: "/usr/local/",
		Err:  errors.New("666"),
	}
	if e, ok := fe.(*os.PathError); ok {
		fmt.Printf("%v\n", e)
	}
	// 或者
	/*switch err := err.(type) {
	case ParseError:
		PrintParseError(err)
	case PathError:
		PrintPathError(err)
		...
	default:
		fmt.Printf(“Not a special 10error, just %s\n”, err)
	}*/
}

// 计算平方根
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	return math.Sqrt(f), nil
}

type FileError struct {
	Op   string
	Path string
	Err  error
}

func (f *FileError) Error() string {
	return f.Op + ": " + f.Path + ": " + f.Err.Error()
}
