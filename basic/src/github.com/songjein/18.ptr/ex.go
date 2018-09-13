package main

import "fmt"

func main() {
	var numPtr *int // default: nil (nil != 0 in go)

	fmt.Println(numPtr)

	var nPtr *int = new(int) // mem alloc

	fmt.Println(nPtr)

	fmt.Println(*nPtr) // dereference

	*nPtr = 55

	fmt.Println(*nPtr)

	var num = 99
	var ptr *int = &num

	fmt.Println(ptr)
	fmt.Println(&num)

	// can't assign actual address into the ptr variable in go
	// ptr = 0x1231231231
	// there's no ptr ops
	// like ptr++

	callByRef(&num)
	fmt.Println(num)
}

func callByRef(n *int) {
	*n = 2
}
