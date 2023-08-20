package main1

func main1() {
	_ = stackIt1()
}

//go:noinline
func stackIt1() int {
	y := 2
	return y * 2
}

/*
Go compilers will allocate variables that are local to a function in that function’s stack frame.

Try it: (-m: prints optimization decisions)(-l: disable/omit inlining decisions)
$ cd golang/allocation/main1
$ go build -gcflags '-m -l'
>>> you will see no output of optimization decisions...
*/
