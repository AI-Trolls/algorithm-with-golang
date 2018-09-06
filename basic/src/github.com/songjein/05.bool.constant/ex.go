package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false

	const age int = 10
	const name string = "abc"
	// const score int --> compile error

	const a = 10

	fmt.Println(b1, b2, age, name, a)

	const x, y int = 1, 2
	const aa, bb = 10, "abc"

	fmt.Println(x, y, aa, bb)

	const (
		xxx, yyy int = 1, 2
		aaa, bbb     = 10, "aaa"
	)

	// enum
	const (
		Sunday = iota // 0, if iota + 1 --> sunday : 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Thursday)

	const (
		q = 1 << iota
		w = 1 << iota
		e = 1 << iota
		r = 1 << iota
	)
	fmt.Println(q, w, e, r)

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1
		bit1, mask1
		_, _
		bit3, mask3
	)
	fmt.Println(bit3, mask3)
}
