package main

import (
	"fmt"
	"math"
)

func main() {

	var num int = 1

	var fnum float32 = 0.1
	var fnum2 float32 = .12
	var fnum3 float64 = .12345e+2
	var fnum4 float64 = 5.12345e-5

	fmt.Println(num, fnum, fnum2, fnum3, fnum4)

	// Rounding Error
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)

	const eps = 1e-14 // machine epsilon for go
	fmt.Println(math.Abs(a-9.0) <= eps)

	// byte
	var bnum byte = 10
	var bnum2 byte = 0x32 // 16 base
	var bnum3 byte = 'a'  // can't use "double qoute "

	// rune for unicode(UTF-8)
	// var r1 rune = ''
}
