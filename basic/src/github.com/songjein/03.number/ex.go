package main

import (
	"fmt"
	"math"
	"unsafe"
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

	// rune for unicode(UTF-8), can't use ""
	var r1 rune = '한'
	var r2 rune = '\ud55c'
	var r3 rune = '\U0000d55c'

	fmt.Println(bnum, bnum2, bnum3, r1, r2, r3)

	// ops for num
	// <<, >>, ^(bit flipping)
	// 1 + 1.2 --> error (both operand must have the same type)

	// max, min value
	var n1 uint8 = math.MaxUint8   // + 1 --> compile error!! (overflow)
	var n2 uint16 = math.MaxUint16 // -1 --> compile error (underflow --> overflow)
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	fmt.Println(unsafe.Sizeof(n1), unsafe.Sizeof(n2)) // byte size
}
