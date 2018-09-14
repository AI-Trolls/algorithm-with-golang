package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("quack")
}

func (d Duck) feathers() {
	fmt.Println("duck has white and gray feathers")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("quaaaack")
}

func (p Person) feathers() {
	fmt.Println("person has non feathers")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var john Person

	inTheForest(donald)
	inTheForest(john)

	// and we can check whether some types implement certain interface
	if v, ok := interface{}(donald).(Quacker); ok {
		// v -> donald
		// ok -> true / false
		fmt.Println(v, ok)
	}
}
