# optionalarg
Optionalarg provides a convenient way to pass optional arguments. It
is a light-weight way to define options.

The idea is define a distinct type for each parameter, and use type
reflection to match parameters with arguments.

For example, to pass two optional string options to a function `Foo`,
define two types for the options,
```
type Name string
type Value string
```
and pass the options to `Foo`,
```
Foo(Value("world"), Name("world"))
```
`Foo` uses `optionalarg.Set` to set the arguments,
```
func Foo(args ...interface{}) {
	var name, value string
	optionalarg.Set(args, (*Name)(&name), (*Value)(&value))
	...
}
```
