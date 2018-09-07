package main

import "fmt"

/*
 *	closure ; close over
 *	closure can access the outer function's variable
 **/
func main() {
	a, b := 3, 5

	f := func(x int) int {
		return a*x + b
	}

	y := f(5)

	fmt.Println(y)

	f2 := calc()
	fmt.Println(f2(1))
	fmt.Println(f2(2))
}

/*
 *	can save the env/status of the function
 **/
func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}
