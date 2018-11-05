package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU()) // utilize all of the cpus in this system
	runtime.GOMAXPROCS(1) // utilize all of the cpus in this system

	fmt.Println(runtime.GOMAXPROCS(0)) // pass 0 : print current status

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i) /* param 'i' important!! */
	}

	/*
		closure executed by goroutine (in loop) -> executed after the loop termination
			fm.Println(s, i) -> 100 100 100 100 100 100 ...

		therefore, variables which are modified by the loop must be passed to the params
	*/

	fmt.Scanln()
}
