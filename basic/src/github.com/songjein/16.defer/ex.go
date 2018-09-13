package main

import (
	"fmt"
	"os"
)

func main() {
	// defer -> LIFO (called right before this function's end)
	defer world()

	hello()
	hello()
	hello()

	// defer -> LIFO ; useful case : closing file
	defer func() {
		fmt.Println("deeeefeeeeeeeeeeeeeeeeeeer")
	}()

	ReadHello()

}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
