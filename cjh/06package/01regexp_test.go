package _package

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	// 简单模式
	// ok, err := regexp.Match("\\d{2}", []byte("2012-12-12 12:21:12"))
	ok, err := regexp.MatchString("\\d{2}", "2012-12-12 12:21:12")
	fmt.Printf("%t\n", ok)
	if !ok {
		fmt.Printf("%v\n", err)
	}
	// 更多功能要Compile新建一个Regex
	compile, _ := regexp.Compile("\\d{2}")
	str := "12cde"
	str = compile.ReplaceAllString(str, "ab")
	fmt.Printf("%s\n", str) // abcde
}

// Compile 函数也可能返回一个错误，我们在使用时忽略对错误的判断是因为我们确信自己正则表达式是有效的。
// 当用户输入或从数据中获取正则表达式的时候，我们有必要去检验它的正确性。
// 另外我们也可以使用 MustCompile 方法，它可以像 Compile 方法一样检验正则的有效性，但是当正则不合法时程序将 panic
