package main

import "fmt"

// interface just define a set of definitions of methods

type hello interface {
}

type MyInt int // also possible

func (i MyInt) Print() {
	fmt.Println(i)
}

////////////////////////////////////////

type Printer interface {
	Print()
}

////////////////////////////////////////

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

func main() {
	// ex 1
	var h hello
	fmt.Println(h)

	// ex 2
	var i MyInt = 5
	var p Printer // interface use -er postfix often
	p = i
	p.Print()

	// ex3
	r := Rectangle{10, 20}
	p = r
	p.Print()

	// ex4
	p1 := Printer(i) // declaration with initialization
	p2 := Printer(r)
	p1.Print()
	p2.Print()

	// ex5
	pArr := []Printer{i, r} // interface slice
	for idx, _ := range pArr {
		pArr[idx].Print()
	}
	for _, value := range pArr {
		value.Print()
	}

	// call 2 different method using the same interface

	// when there are two types..
	// go interface considers both have the same type if they have the same set of methods
	// regardless of their types
}
