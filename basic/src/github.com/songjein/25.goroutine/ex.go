package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	goroutine vs. main func -> executed simultaneously
	- fmt.Println cannot be executed (main func -> returned earlier than fmt.Println)
	-> use fmt.Scanln() (wait until press enter)

	time.Duration -> nano seconds (1 sec : 1000000000 nano sec)
	cf) 1000000 micro sec -> 1 sec

	10 sec : time.Second * 10
	35 micro sec : time.Microsecond * 35
	120 milli sec : time.Millisecond * 120
*/

func hello(n int) {
	r := rand.Intn(100) // random waiting time
	time.Sleep(time.Duration(r) /*type casting*/)
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // run func as goroutine
	}

	fmt.Scanln() // prevent the termination of main
}
