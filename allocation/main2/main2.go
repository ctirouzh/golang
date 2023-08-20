package main2

func main2() {
	_ = stackIt2()
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

/*
Go compilers will allocate variables that are local to a function in that function’s stack frame.
However, if the compiler cannot prove that the variable is not referenced after the function returns,
then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.

Try it: (-m: prints optimization decisions)(-l: disable/omit inlining decisions)

$ cd golang/allocation/main2
$ go build -gcflags '-m -l'
>>> # github.com/foldera/golang/allocation/main2
>>> ./main2.go:10:2: moved to heap: res

*/
