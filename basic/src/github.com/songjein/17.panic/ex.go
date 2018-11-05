package main

import "fmt"

/**
 *	panic ; error -> program termination
 *	ex) index out of range
 *	user can invoke the error using panic() function
 */
func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	panic("Custom Error!")
	fmt.Println("vim-go")
}

func f2() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()
	f2()

	fmt.Println("Hello, world")
}
