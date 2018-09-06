package main

import "fmt" // alias --> import f "fmt"
// import . "fmt" --> get "fmt" as a global package

func main() {
	fmt.Println(1 == 1)
	fmt.Println(3.5 == 3.5)
	fmt.Println("hello" == "hello")

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b) // check each elems equality

	c := []int{1, 2, 3}
	fmt.Println(c == nil) // check the mem is allocated

	e := 1
	var p1 *int = &e
	var p2 *int = &e
	fmt.Println(p1 == p2)
}
