package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	// 字符串是一种值类型，其值不可变。更深入的讲：字符串是定长字节数组
	// 转义字符：\U \u Unicode字符
	// 解释字符串，用双引号括起来，这些字符会被转义
	a := "a\nb"
	fmt.Println(a)
	// 非解释字符串，用反单引号括起来，其中是什么就打印什么
	fmt.Println(`a\nb`)

	//字符串的零值是空字符串""

	//字符串的内容（纯字节）可以通过标准索引来获取，索引从0开始
	b := "abcdefg"
	fmt.Println(b[0])        //a
	fmt.Println(b[len(b)-1]) //g
	// 这种转换方案只对纯 ASCII 码的字符串有效

	// 字符串拼接用+，也可用+=
	c := "Hel" + "lo, "
	c += "World!"
	fmt.Println(c)

	//+并不是最有效率的拼接方式，可以使用strings.join()
	s1 := "a,b,c"
	split := strings.Split(s1, ",")
	join := strings.Join(split, "|")
	fmt.Println(join)

	//最好的方式是使用bytes.Buffer,这种方式比+=更节省CPU和内存
	var buffer bytes.Buffer
	buffer.WriteString(join)
	buffer.WriteString("|d")
	fmt.Println(buffer.String())
}

func TestTwo(t *testing.T) {
	//判断前缀
	println(strings.HasPrefix("abc", "a")) //true
	//判断后缀
	println(strings.HasSuffix("abc", "e")) //false
	//包含
	println(strings.Contains("abc", "a"))
	println(strings.ContainsAny("abc", "a"))
	//判断索引位置-1表示不存在
	println(strings.Index("abcd", "bc"))
	println(strings.Index("abcd", "e"))
	// LastIndex最后出现的索引
	println(strings.LastIndex("aba", "a"))
	//对于非ASCII采用Rune
	println(strings.IndexRune("中华人民共和国", 3))
}
