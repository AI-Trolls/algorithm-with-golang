package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

type Person struct {
	name       string
	age, money int
}

func main() {
	var p Person // local variable, each fields have default values
	fmt.Println(p)

	var rect *Rectangle = new(Rectangle)
	rect2 := new(Rectangle)
	fmt.Println(rect, rect2)

	// declare with init
	var rect3 Rectangle = Rectangle{10, 20}
	rect4 := Rectangle{20, 10}
	rect5 := Rectangle{width: 90, height: 100}
	fmt.Println(rect3, rect4, rect5)

	// access fields of struct
	var per1 Person
	var per2 *Person = new(Person)
	per1.name = "jein" // both types have the same way
	per2.name = "jem"
	fmt.Println(per1)
	fmt.Println(per2)

	// there's no way to malloc a struct and init that one simultaneously
	// but we can use the 'constructor pattern'
	// golang can return 'the struct or ptr of struct' which are declared as local variables in the function
	NewRectangle := func(width, height int) *Rectangle {
		return &Rectangle{width, height}
	}
	rrrr := NewRectangle(100, 200)
	fmt.Println(rrrr)

	rrrr2 := &Rectangle{20, 10} // same with the above method
	fmt.Println(rrrr2)

	rectangleArea := func(rect *Rectangle) int { // if not use ptr for struct, all of the field values copied when the params are passed
		return rect.width * rect.height
	}

	fmt.Println(rectangleArea(&rect3))
	fmt.Println(rectangleArea(rrrr))
}
