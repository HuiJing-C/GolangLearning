package _package

import (
	. "github.com/HuiJing-C/GolangLearning/cjh/06package/pack" // cjh前的内容为go.mod的module名
	"testing"
)

// 主程序利用的包必须在主程序编写之前被编译。
// 导入外部安装包:
// 如果你要在你的应用中使用一个或多个外部包，首先你必须使用 go install address 在你的本地机器上安装它们
// 假设你想使用 http://codesite.ext/author/goExample/goex 这种托管在 Google Code、GitHub 和 Launchpad 等代码网站上的包。
// 你可以通过如下命令安装：
// go install codesite.ext/author/goExample/goex
// 将一个名为 codesite.ext/author/goExample/goex 的 map 安装在 $GOROOT/src/ 目录下。
// 通过以下方式，一次性安装，并导入到你的代码中：
// import goex "codesite.ext/author/goExample/goex"
// 因此该包的 URL 将用作导入路径。
// 在 http://golang.org/cmd/goinstall/ 的 go install 文档中列出了一些广泛被使用的托管在网络代码仓库的包的导入路径

func TestPackage(t *testing.T) {
	// 当使用.来做为包的别名时，你可以不通过包名来使用其中的项目
	// 当使用_来做包名时，只导入其副作用，也就是说，只执行它的init函数并初始化其中的全局变量。
	println(PackInt)
	println(ReturnStr())
}

// 包的初始化:
// 程序的执行开始于导入包，初始化 main 包然后调用 main 函数。
// 一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 init 函数来初始化。一个包可能有多个 init 函数甚至在一个源码文件中。它们的执行是无序的。这是最好的例子来测定包的值是否只依赖于相同包下的其他值或者函数。
// init 函数是不能被调用的。
// 导入的包在包自身初始化前被初始化，而一个包在程序执行中只能初始化一次。
