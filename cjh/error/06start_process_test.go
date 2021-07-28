package error

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

/*
os 包有一个 StartProcess 函数可以调用或启动外部系统命令和二进制可执行文件；
它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
这个函数返回被启动进程的 id（pid），或者启动失败返回错误。
exec 包中也有同样功能的更简单的结构体和函数；主要是 exec.Command(name string, arg ...string) 和 Run()。
首先需要用系统命令或可执行文件的名字创建一个 Command 对象，然后用这个对象作为接收者调用 Run()。
下面的程序（因为是执行 Linux 命令，只能在 Linux 下面运行）演示了它们的使用：
*/

func TestProcessOne(t *testing.T) {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)

	println("=========================")

	// 2nd example: show all processes
	pid, err = os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}

func TestProcessTwo(t *testing.T) {
	// 2) exec.Run //
	/***************/
	// Linux:  OK, but not for ls ?
	// cmd := exec.Command("ls", "-l")  // no error, but doesn't show anything ?
	// cmd := exec.Command("ls")        // no error, but doesn't show anything ?
	cmd := exec.Command("gedit") // this opens a gedit-window
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v\n", cmd)
	// The command is &{/bin/ls [ls -l] []  <nil> <nil> <nil> 0xf840000210 <nil> true [0xf84000ea50 0xf84000e9f0 0xf84000e9c0] [0xf84000ea50 0xf84000e9f0 0xf84000e9c0] [] [] 0xf8400128c0}
}

// in Windows: uitvoering: Error fork/exec /bin/ls: The system cannot find the path specified. starting process!
