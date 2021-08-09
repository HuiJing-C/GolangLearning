package ArrayAndSlice

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
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

	// 在内存中，字符串是一个双字结构，一个是指向实际地址的指针，一个记录字符串长度的整数
}

func TestChangeString(t *testing.T) {
	// 字符串是不可变的，不可以通过str[index]修改字符串内容，要先转换字节数组，然后改变字节数组的值然后再转为str
	s := "12345"
	// s[0] = '2' // cannot assign to s[0]
	bytes := []byte(s)
	bytes[0] = '2'
	s2 := string(bytes)
	fmt.Printf("%s\n", s2) // 22345
}

func TestSort(t *testing.T) {
	// sort包下有搜索和排序操作
	ints := []int{2, 3, 4, 1}
	sort.Ints(ints)
	fmt.Printf("%v\n", ints) // [1 2 3 4]
	// 检查是否一排序
	fmt.Printf("%t\n", sort.IntsAreSorted(ints)) // true
	// 想要在数组或切片中搜索一个元素，该数组或切片必须先被排序（因为标准库的搜索算法使用的是二分法）。
	// 然后，就可以使用函数 func SearchInts(a []int, n int) int 进行搜索，并返回对应结果的索引值。
	// 官方文档 http://golang.org/pkg/sort/
}

func TestAppend(t *testing.T) {
	// append有很多操作
	// 1 将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
	//
	// 2 复制切片 a 的元素到新的切片 b 上：
	//
	// b = make([]T, len(a))
	// copy(b, a)
	// 3 删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
	//
	// 4 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
	//
	// 5 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
	//
	// 6 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
	//
	// 7 在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
	//
	// 8 在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
	//
	// 9 取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]
	//
	// 10 将元素 x 追加到切片 a：a = append(a, x)
	// 从数学的角度来看，切片相当于向量，如果需要的话可以定义一个向量作为切片的别名来进行操作。
	// 如果需要更加完整的方案，可以学习一下 Eleanor McHugh 编写的几个包：slices(http://github.com/feyeleanor/slices)、chain(http://github.com/feyeleanor/chain) 和 lists(http://github.com/feyeleanor/lists)。
}

func TestGC(t *testing.T) {
	// 切片和垃圾回收
	// 切片的底层指向一个数组，该数组的实际体积可能要大于切片所定义的体积。
	// 只有在没有任何切片指向的时候，底层的数组内层才会被释放，这种特性有时会导致程序占用多余的内存。

	// 参考FindDigits1
	// 这段代码可以顺利运行，但返回的 []byte 指向的底层是整个文件的数据。只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。
	// 换句话说，一点点有用的数据却占用了整个文件的内存。
	// 想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片中：参考FindDigits2
}
func FindDigits1(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func FindDigits2(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
