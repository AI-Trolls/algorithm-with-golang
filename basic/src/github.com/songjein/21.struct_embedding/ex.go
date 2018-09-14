package main

import "fmt"

// there's no inheritance in go
// but can have the same effact by using embedding

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	p      Person // has-a
	school string
	grade  int
}

type Student2 struct {
	Person // is-a
	school string
	grade  int
}

// overriding
func (p *Student2) greeting() {
	fmt.Println("wow")
}

func main() {

	var s Student
	s.p.greeting()

	var s2 Student2
	s2.greeting()
	s2.Person.greeting()

}

// multiple embedding -> ok -> but, better to use interface in go
