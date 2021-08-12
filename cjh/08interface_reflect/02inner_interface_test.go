package _8interface_reflect

import "testing"

// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
type A interface {
	read() string
}

type B interface {
	write(s string)
}

type C interface {
	A
	B
	close()
}

type D struct{}

func (d *D) read() string {
	return "D read"
}

func (d *D) write(str string) {}
func (d *D) close()           {}

func TestInnerInterface(t *testing.T) {
	var c C
	c = &D{}
	println(c.read()) // D read
}
