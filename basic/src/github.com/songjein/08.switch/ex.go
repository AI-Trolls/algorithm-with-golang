package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := 2

	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(-1)
	}

	s := "world"

	switch s {
	case "hello":
		fmt.Println("h")
	case "world":
		fmt.Println("w")
	default:
		fmt.Println("no")
	}

	ii := 9
	ss := "hello"

	switch ii {
	case 1:
		fmt.Println(10)
	case 2:
		if ss == "hello" {
			fmt.Println("break")
			break
		}
		fmt.Println(20)
		fallthrough
	case 3:
		fmt.Println(300)
	case 4, 5, 6, 7:
		fmt.Println(4567)
	}

	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(10); {
	case i >= 3 && i < 6:
		fmt.Println("36")
	case i == 9:
		fmt.Println("9")
	default:
		fmt.Println(i)
	}
}
