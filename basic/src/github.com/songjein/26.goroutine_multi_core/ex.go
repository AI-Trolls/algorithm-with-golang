package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // utilize all of the cpus in this system

	fmt.Println(runtime.GOMAXPROCS(0)) // pass 0 : print current status

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
