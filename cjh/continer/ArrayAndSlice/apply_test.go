package ArrayAndSlice

import (
	"fmt"
	"testing"
)

func TestStringToSlice(t *testing.T) {
	// 从字符串生成字节切片
	// 字符串本质上是一个字节数组，可以通过以下方式创建切片
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c) // 0:ÿ 2:界
	}
	slice1 := []byte(s)
	fmt.Printf("%d %d\n", len(slice1), cap(slice1)) // 5 32
	for i, b := range slice1 {
		fmt.Printf("%d %b\n", i, b)
	}
	// 0 11000011
	// 1 10111111
	// 2 11100111
	// 3 10010101
	// 4 10001100

	// 还可以使用copy()进行此操作
	var dst []byte = make([]byte, 5)
	copy(dst, "123")
	fmt.Printf("%d %d\n", len(dst), cap(dst)) // 5 5
	for i, b := range dst {
		fmt.Printf("%d %b\n", i, b)
	}
	// 0 110001
	// 1 110010
	// 2 110011
	// 3 0
	// 4 0

	// 和字符串转换一样，同样可以使用 c := []int(s) 语法，这样切片中的每个 int 都会包含对应的 Unicode 代码，因为字符串中的每次字符都会对应一个整数。
	// 类似的，也可以将字符串转换为元素类型为 rune 的切片：r := []rune(s)。
	// 可以通过代码 len([]int(s)) 来获得字符串中字符的数量，但使用 utf8.RuneCountInString(s) 效率会更高一点。

	// 还可以将一个字符串追加到某一个字符数组的尾部
	var b []byte
	fmt.Printf("%d\n", len(b)) // 0
	b = append(b, s...)
	fmt.Printf("%d\n", len(b)) // 5
}

func TestSubString(t *testing.T) {
	// 使用 substr := str[start:end] 可以从字符串 str 获取到从索引 start 开始到 end-1 位置的子字符串。
	// 同样的，str[start:] 则表示获取从 start 开始到 len(str)-1 位置的子字符串。而 str[:end] 表示获取从 0 开始到 end-1的子字符串。
	s := "abcdefg"
	s1 := s[0:3]
	s2 := s[3:]
	s3 := s[:7]
	fmt.Printf("%s %s %s %s\n", s, s1, s2, s3) // abcdefg abc defg abcdefg
}
