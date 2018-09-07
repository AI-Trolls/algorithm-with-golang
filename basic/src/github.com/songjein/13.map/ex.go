package main

import "fmt"

func main() {
	/*
	 *	declaration
	 **/
	var a map[string]int = make(map[string]int)

	var b = make(map[string]int)

	fmt.Println(a, b)

	c := make(map[string]int)
	c["count"] = 100
	c["ggggg"] = 200

	fmt.Println(c)

	fmt.Println("print non-exist :", c["aaaaa"]) // empty value int: 0, float: 0.0, string: ""

	aa := map[string]int{"abc": 10, "def": 20}

	bb := map[string]int{
		"aa": 11,
		"bb": 22,
	}

	fmt.Println(aa, bb)

	/*
	 *	check existance
	 **/
	if value, ok := bb["aa"]; ok {
		fmt.Println(value)
	}

	/*
	 * map interation
	 **/
	for key, value := range bb {
		fmt.Println("iter", key, value)
	}

	for _, value := range bb {
		fmt.Println("iter without key", value)
	}

	/*
	 *	delete
	 **/
	abc := map[string]int{"Hello": 10, "World": 20}
	delete(abc, "World")
	fmt.Println("deleted result", abc)

	/*
	 *	map in map
	 **/
	mim := map[string]map[string]float32{
		"AAA": map[string]float32{
			"aaa": 111,
		},
		"BBB": map[string]float32{
			"bbb": 222,
		},
		"CCC": map[string]float32{
			"ccc": 333,
		},
	}

	fmt.Println(mim)

}
