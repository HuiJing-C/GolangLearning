package _8interface_reflect

import "testing"

type K interface {
	read()
}
type K2 interface {
	read()
}
type L struct{}

func (l *L) read() {
	println("L read")
}

func (l *L) write() {
	println("L write")
}

func TestImplement(t *testing.T) {
	var k K
	k = &L{}
	if kv, ok := k.(K); ok {
		kv.read()
	}
}
