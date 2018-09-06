package main

import "fmt"

func main() {
	// 선언법
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	var c = [5]int{2, 3, 4, 5, 6}
	d := [5]int{5, 4, 3, 2, 1}

	fmt.Println(b, c, d)

	e := [5]int{
		4,
		5,
		6,
		7,
	}

	f := [...]int{2, 2, 2, 2, 2} // auto len

	g := [...]string{"aaa", "bbb", "ccc"}

	fmt.Println(e, f, g)

	arr2d := [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(arr2d)

	// 루프 순회

	aaa := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(aaa); i++ {
		fmt.Println(aaa[i])
	}

	// for _, value := range aaa {
	for i, value := range aaa {
		fmt.Println(i, value)
	}

	bbb := aaa // copy 'aaa' is not a pointer
	bbb[2] = 2222
	fmt.Println(aaa, bbb)

}
