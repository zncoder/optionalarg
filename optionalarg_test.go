package optionalarg_test

import (
	"testing"
	"time"

	"github.com/zncoder/optionalarg"
)

func TestOptionalArg(t *testing.T) {
	type Count int
	type Name string
	type Deadline time.Time

	foo := func(args ...interface{}) (int, string, time.Time) {
		var i int
		var s string
		var deadline time.Time
		optionalarg.Set(args, (*Count)(&i), (*Name)(&s), (*Deadline)(&deadline))
		return i, s, deadline
	}

	t.Run("set arguments", func(t *testing.T) {
		c, n, d := foo(Count(1), Name("a"))
		if c != 1 {
			t.Error("c should be 1")
		}
		if n != "a" {
			t.Error("n should be a")
		}
		if !d.IsZero() {
			t.Error("d should be zero")
		}
	})

	t.Run("duplicate arg type panics", func(t *testing.T) {
		defer func() {
			if v := recover(); v == nil || v != "type of 2-th arg:optionalarg_test.Count is dup" {
				t.Error("Set should panic with: arg...Count is dup")
			}
		}()

		foo(Count(1), Name("a"), Count(2))
	})

	t.Run("unknown arg type panics", func(t *testing.T) {
		defer func() {
			if v := recover(); v == nil || v != "0-th arg:int/1 has no matching param type" {
				t.Errorf("Set should panic with: ...no matching param type, v:%q", v)
			}
		}()

		foo(1, Name("a"), Count(2))
	})

	t.Run("duplicate dest type panic", func(t *testing.T) {
		bar := func(args ...interface{}) {
			var i, j int
			var s string
			optionalarg.Set(args, (*Count)(&i), (*Name)(&s), (*Count)(&j))
		}

		defer func() {
			if v := recover(); v == nil || v != "type of 2-th dest:*optionalarg_test.Count is dup" {
				t.Errorf("Set should panic with: dest...Count is dup, v:%q", v)
			}
		}()

		bar(Count(1), Name("a"))
	})
}
