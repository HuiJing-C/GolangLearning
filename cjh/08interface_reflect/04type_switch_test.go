package _8interface_reflect

import (
	"fmt"
	"testing"
)

// 类型判断：type-switch
// type-switch 不允许有 fallthrough
type H interface {
	read()
}

type I struct{}
type J struct{}

func (this *I) read() {
	println("I read")
}
func (this *J) read() {
	println("J read")
}
func TestTypeSwitch(t *testing.T) {
	tps := []interface{}{&I{}, &J{}, nil, "123"}
	typeSwitch(tps)
}

func typeSwitch(types []interface{}) {
	for i := range types {
		switch t := types[i].(type) {
		case *I:
			t.read()
		case *J:
			t.read()
		case nil:
			println("Nil")
		default:
			fmt.Printf("Unexpected Type %v\n", types[i])
		}
	}
}
