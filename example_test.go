package optionalarg_test

import (
	"fmt"
	"time"

	"github.com/zncoder/optionalarg"
)

type (
	Name     string
	Duration time.Duration
	Value    string
)

func Foo(msg string, args ...interface{}) {
	var name, value string
	var duration time.Duration
	optionalarg.Set(args, (*Name)(&name), (*Value)(&value), (*Duration)(&duration))
	fmt.Printf("msg:%s name:%s value:%s duration:%v\n", msg, name, value, duration)
}

func Example() {
	Foo("hello", Name("foo"), Duration(time.Second))
	// Output: msg:hello name:foo value: duration:1s
}
