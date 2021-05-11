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

func TestThree(t *testing.T) {
	// strings.Replace("str", "old", "new", n),将str中前n个old字符串替换为new,n=-1时替换所有
	println(strings.Replace("abcdaefg", "a", "b", -1)) //bbcdbefg

	// Count("str", "s")计算str字符串中s出现的非重叠次数
	println(strings.Count("ggggg", "gg")) //2

	// strings.Repeat("str", n) 返回一个字符串，由str重复n次
	println(strings.Repeat("str", 3)) //strstrstr
	//大小写转换，将字符串中所有Unicode字符全部转换为对应的大小写
	println(strings.ToUpper("aBc")) // ABC
	println(strings.ToLower("A曹b")) // a曹b

	//TrimSpace(s)剔除s开头和结尾的空白符号,Trim(str,s)替换掉str开头和结尾的s,TrimLeft/TrimRight提出开头或者结尾的字符串
	println(strings.TrimSpace("\t\n abc \t\\"))       //abc 	\
	println(strings.Trim("\tnb\tnnjs\tn", "\tn"))     //b	nnjs
	println(strings.TrimLeft("\tnb\tnnjs\tn", "\tn")) //b	nnjs	n

	// strings.Fields(s)将s以动态的空白字符串分割为字符串slice,如果s全是空白符，则返回长度为零的slice
	// strings.Split(s, sep)自定义分隔符，同样返回slice
	fmt.Printf("%v\n", strings.Fields("\na b\t\nc")) //[a b c]

	// strings.join(sl []string, sep string),将s1切片中的所有元素用分隔符sep拼接为一个长字符串
	println(strings.Join(strings.Fields("\na b\t\nc"), "|")) //a|b|c
}

func TestFour(t *testing.T) {
	// 读取字符串中的内容
	reader := strings.NewReader("abc")
	println(reader.Size())
}
