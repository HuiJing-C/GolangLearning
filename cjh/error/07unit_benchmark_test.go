package error

import "testing"

/*
首先所有的包都应该有一定的必要文档，然后同样重要的是对包的测试。
Go 的测试工具 gotest，我们会用更多的例子进行详细说明。
名为 testing 的包被专门用来进行自动化测试，日志和错误报告。并且还包含一些基准测试函数的功能。
备注：gotest 是 Unix bash 脚本，所以在 Windows 下你需要配置 MINGW 环境试代码和；在 Windows 环境下把所有的 pkg/linux_amd64 替换成 pkg/windows。
对一个包做（单元）测试，需要写一些可以频繁（每次更新后）执行的小块测试单元来检查代码的正确性。于是我们必须写一些 Go 源文件来测试代码。
测试程序必须属于被测试的包，并且文件名满足这种形式 *_test.go，所以测包中的业务代码是分开的。
*_test 程序不会被普通的 Go 编译器编译,其中的方法和属性也不会被导出，所以当放应用部署到生产环境时它们不会被部署；只有 gotest 会编译所有的程序：普通程序和测试程序。
测试文件中必须导入 "testing" 包，并写一些名字以 TestZzz 打头的全局函数（不然不会跑），这里的 Zzz 是被测试函数的字母描述，如 TestOne，TestTwo 等。
测试函数必须有这种形式的头部：
func TestA(t *testing.T)

T 是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如 t.Log，t.Error，t.ErrorF 等。
在函数的结尾把输出跟想要的结果对比，如果不等就打印一个错误。成功的测试则直接返回。
用下面这些函数来通知测试失败：

1）func (t *T) Fail()	标记测试函数为失败，然后继续执行（剩下的测试）。
2）func (t *T) FailNow()	标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件。
3）func (t *T) Log(args ...interface{})	args 被用默认的格式格式化并打印到错误日志中。
4）func (t *T) Fatal(args ...interface{})	结合 先执行 3），然后执行 2）的效果。

由于测试需要具体的输入用例且不可能测试到所有的用例（非常像一个无穷的数），所以我们必须对要使用的测试用例思考再三。
至少应该包括：
---正常的用例
---反面的用例（错误的输入，如用负数或字母代替数字，没有输入等）
---边界检查用例（如果参数的取值范围是 0 到 1000，检查 0 和 1000 的情况）

运行 go test 来编译测试程序，并执行程序中所有的 TestZZZ 函数。如果所有的测试都通过会打印出 PASS。
gotest 可以接收一个或多个函数程序作为参数，并指定一些选项。
结合 --chatty 或 -v 选项，每个执行的测试函数以及测试状态会被打印。

go test 00first_test.go -v
=== RUN   TestOne
--- PASS: TestOne (0.00s)
=== RUN   TestTwo2
--- PASS: TestTwo2 (0.00s)
PASS
ok      command-line-arguments  0.078s
*/

/*
testing 包中有一些类型和函数可以用来做简单的基准测试；测试代码中必须包含以 BenchmarkZzz 打头的函数并接收一个*testing.B 类型的参数，比如：
func BenchmarkReverse(b *testing.B) {
    ...
}
命令 go test –test.bench=.* 会运行所有的基准测试函数；代码中的函数会被调用 N 次（N是非常大的数，如 N = 1000000），
并展示 N 的值和函数执行的平均时间，单位为 ns（纳秒，ns/op）。如果是用 testing.Benchmark 调用这些函数，直接运行程序即可。
*/
func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(i)
	}
}

// BenchmarkOne-8   	  125860	      7989 ns/op
