/*
Package optionalarg provides a convenient way to pass optional
arguments. It provides a light-weight way to define options.

The idea is define a distinct type for each parameter, and use type
reflection to match parameters with arguments.

For example, to pass two optional string options to a function Foo,
define two types for the options,

	type Name string
	type Value string

and pass the options to Foo,

	Foo(Value("world"), Name("world"))

Foo uses optionalarg.Set to set the arguments,

	func Foo(args ...interface{}) {
		var name, value string
		optionalarg.Set(args, (*Name)(&name), (*Value)(&value))
		...
	}

Tip: Group these types for nice godoc formatting.
*/
package optionalarg

import (
	"fmt"
	"reflect"
)

// Set sets arguments to dests, based on the type of dests.
func Set(args []interface{}, dests ...interface{}) {
	// set up expected types from dests
	m := make(map[reflect.Type]reflect.Value, len(dests))
	for i, x := range dests {
		t := reflect.TypeOf(x)
		if t.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("%d-th dest:%T is not pointer", i, x))
		}

		v := reflect.ValueOf(x).Elem()
		if !v.CanSet() {
			panic(fmt.Sprintf("%d-th dest:%T is not setable", i, x))
		}

		if _, ok := m[t.Elem()]; ok {
			panic(fmt.Sprintf("type of %d-th dest:%T is dup", i, x))
		}
		m[t.Elem()] = v
	}

	// set args to dests
	set := make(map[reflect.Type]struct{})
	for i, x := range args {
		t := reflect.TypeOf(x)
		if _, ok := set[t]; ok {
			panic(fmt.Sprintf("type of %d-th arg:%T is dup", i, x))
		}
		v, ok := m[t]
		if !ok {
			panic(fmt.Sprintf("%d-th arg:%T/%v has no matching param type", i, x, x))
		}
		v.Set(reflect.ValueOf(x))
		set[t] = struct{}{}
	}
}
