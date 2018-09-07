package main

import "fmt"

// func name(param type) return-type {}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	r = sum2(5, 5)
	fmt.Println(r)

	s, m := sumAndMul(3, 5)
	fmt.Println(s, m)

	s, m = sumAndMul2(3, 5)
	fmt.Println(s, m)

	s = sum3(1, 2, 3, 4, 5)
	fmt.Println(s)

	/*
	 *	Variable length arguments
	 **/
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a, 5, 5)
	s = sum3(a...)
	fmt.Println(s)

	fmt.Println(factorial(5))

	/*
	 *	save function as a variable
	 **/
	var sumFunc func(a int, b int) int = sum
	// or
	sumFunc2 := sum

	fmt.Println(sumFunc(1, 2))
	fmt.Println(sumFunc2(5, 2))

	// slice for function
	f := []func(int, int) int{sum, sum2}
	fmt.Println(f[0](22, 22))
	fmt.Println(f[1](22, 22))

	// map for function
	fm := map[string]func(int, int) int{
		"sum":  sum,
		"sum2": sum2,
	}
	fmt.Println(fm["sum"](111, 111))

	// ananymous function
	func() {
		fmt.Println("Ananymous func")
	}()

	func(s string) {
		fmt.Println(s)
	}("hi")

	rr := func(a int, b int) int {
		return a + b
	}(55, 55)

	fmt.Println(rr)

}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (r int) {
	r = a + b
	return
}

func sumAndMul(a int, b int) (int, int) {
	return a + b, a * b
}

func sumAndMul2(a int, b int) (s int, m int) {
	s = a + b
	m = a * b
	return
}

// variable length arguments (slice)
func sum3(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
