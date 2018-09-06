package main

import "fmt"

func main() {
	if i := 10; i > 5 {
		fmt.Println("over 5")
	} else {
		fmt.Println("under 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 2
	for i < 5 {
		fmt.Println("a")
		i++
	}

	// inf loop for {}

	// using label
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue Loop2
			}

			fmt.Println(i, j)
		}
	}
	fmt.Println("hi")

	// multiple variable --> have to use parallel assignment
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// goto LABEL is possible

	var a int = 1

	if a == 1 {
		goto ERROR1
	}

	if a == 2 {
		goto ERROR2
	}

	fmt.Println("Normal Exit")
	return

ERROR1:
	fmt.Println("Error 1")
	return
ERROR2:
	fmt.Println("Error 2")
	return
}
