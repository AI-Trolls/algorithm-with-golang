package main

import (
	"fmt"
	"strconv"
)

func f1(arg interface{}) { // empty interface can receive all of the types
}

type Any interface{}

func f2(arg Any) {
}

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) { // get the type saved in interface (valid only in switch)
	case int:
		i := arg.(int) // 'type assertion', we can't just use the empty interface variable, have to cast type
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello world"))

	var jein = new(Person)
	jein.name = "jein"
	jein.age = 51

	fmt.Println(jein) // ptr

	fmt.Println(formatString(jein))

	// so the way to check the type saved in the interface was like below
	var t interface{}
	t = Person{"Maria", 20}

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}
