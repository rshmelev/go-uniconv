package test

import (
	_ "fmt"
	"testing"

	"fmt"

	ser "github.com/rshmelev/go-uniconv/structures"
)

type F struct {
	I int
}

type X struct {
	A string
	P *F
	N F
}

func TestSomething(t *testing.T) {
	cast := ser.Cast
	x := &X{}

	s := cast(
		struct {
			A       string
			P       *F
			N       F
			Another *F
		}{"some", &F{10}, F{30}, &F{333}}).JsonStr()
	println(s)

	asjson := cast(s).AsJsonStruct()
	asjson.Into(x)

	assert(asjson.Failed() != nil, "asjson.failed: %v", asjson.Failed())
	assert(x.A != "some", `x.A != "some"`)
	assert(x.P.I != 10, "x.P.I != 10")
	assert(x.N.I != 30, "x.N.I != 30")

}

func assert(b bool, s string, p ...interface{}) {
	if b {
		fmt.Errorf(s, p...)
	}
}
