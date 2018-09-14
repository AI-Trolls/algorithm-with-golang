package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// define receiver variable between 'func' keyword and 'name' of the function
func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

// but there's two way to pass the struct to the receiver variable

func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	// value just copied
	rect.width = rect.width * factor
	rect.height = rect.height * factor
	// can't change struct outside
}

// if we don't need the receiver variable
func (_ Rectangle) dummy() {
	fmt.Println("dummy func")
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area())

	rect.dummy()
}
