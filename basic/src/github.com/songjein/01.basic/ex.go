package main

import "fmt"

func main() {
	i := 10

	if i >= 5 {
		fmt.Println("over 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("hello")
	}
}
