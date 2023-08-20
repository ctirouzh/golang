package main3

func main3() {
	y := 2
	_ = stackIt3(&y) // pass y down the stack as a pointer
}

//go:noinline
func stackIt3(y *int) *int {
	res := *y * 2
	return &res
}

/*
Go compilers will allocate variables that are local to a function in that function’s stack frame.
However, if the compiler cannot prove that the variable is not referenced after the function returns,
then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.

Try it: (-m: prints optimization decisions)(-l: disable/omit inlining decisions)

$ cd golang/allocation/main3
$ go build -gcflags '-m -l'
>>> # github.com/foldera/golang/allocation/main3
>>> ./main3.go:9:15: y does not escape
>>> ./main3.go:10:2: moved to heap: res
*/

/*
Why do we get this seeming inconsistency?

stackIt2 passes the address of y up the stack to main,where y will be referenced after the stack frame of
stackIt2 has already been freed. The compiler is therefore able to judge that y must be moved to the heap
to remain alive. If it does not do this, we’ll get a nil pointer in main when attempted to reference y.

stackIt3, on the other hand, passes y down the stack, and y isn’t referenced anywhere outside main3.
The compiler is therefore able to judge that y can exist within the stack alone, and does not need to
be allocated to the heap. We won’t be able to produce a nil pointer in any circumstances by referencing y.

A general rule we can infer from this is that
	sharing pointers up the stack results in allocations,
	whereas sharing points down the stack does not.
However, this is not guaranteed, so you’ll still need to verify with gcflags or benchmarks to be sure.
What we can say for sure is that any attempt to reduce allocs/op will involve hunting out wayward pointers.
*/
