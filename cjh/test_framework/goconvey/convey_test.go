package goconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOne(t *testing.T) {
	Convey("cjh", t, func() {
		var a interface{}
		So(a, ShouldBeNil)
	})
}
