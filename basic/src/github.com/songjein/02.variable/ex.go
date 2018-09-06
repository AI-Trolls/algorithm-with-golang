package main

import "fmt"
import _ "time"

func main() {
	// two way (1. using var, 2. omit type)
	var i int
	var s string
	var age int = 1
	var name string = "abc"

	// without type (have to assign)
	var num = 1
	var str = "jein"

	// var address --> error

	// short declaration
	sss := "jein"
	nnn := 10

	fmt.Println(i, s, age, name, num, str, sss, nnn)

	// multiple variables
	var x, y int = 30, 50
	var age2, name2 = 10, "Maria"
	a, b, c, d := 1, 1.2, "abc", false

	fmt.Println(x, y, age2, name2, a, b, c, d)

	// Parallel assignment
	var xx, yy int
	var zz int

	xx, yy, zz = 1, 2, 3

	_ = xx
	_ = yy
	_ = zz

	var (
		x1, y1 int = 1, 2
		z1, a1     = 10, "abc"
	)
}
