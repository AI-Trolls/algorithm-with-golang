package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("vim-go")
	var s1 string = "hello world \n"
	var s2 string = "안녕하세요 \n"
	var s3 string = "\ud55c\uae00"
	var s4 string = "\U0000d55c\U0000ae00"
	var s5 string = "\xed\x95\x9c\xea\xb8\x80"

	fmt.Println(s1, s2, s3, s4, s5)

	// multi-line string
	var s6 string = `안녕하세요
Hello world!`

	fmt.Println(s6)

	s7 := "한글"
	s8 := "Hello"

	// string length
	fmt.Println(len(s7), len(s8))

	// utf string length
	fmt.Println(utf8.RuneCountInString(s7))

	// equal
	fmt.Println(s7 == s7)
	fmt.Println(s7 == s8)

	// concat
	fmt.Println(s7 + s7)

	// get char
	fmt.Println(s8[0])

	// can't modify certain index
}
