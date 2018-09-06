package main

import "fmt"

func main() {
	/*
	 *	선언법
	 **/
	var a []int = make([]int, 5) // 'a' is reference type
	var b = make([]int, 5)
	c := make([]int, 5)

	fmt.Println(a, b, c)

	// 생성과 동시 초기화
	aa := []int{1, 2, 3, 4, 5}
	bb := []int{
		2,
		3,
		4,
		5,
	}

	fmt.Println(aa, bb)

	s := make([]int, 5, 10) // len 5(index로 접근가능), capacity 10(실제 메모리 할당 공간)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	k := make([]int, 5, 10)
	fmt.Println(k[4])
	// fmt.Println(k[5]) // runtime error, index out of range

	/*
	 *	slice 붙이기
	 **/
	k = append(k, 4, 5, 6)
	fmt.Println(k)

	aaa := []int{1, 2, 3}
	bbb := []int{4, 5, 6}

	ccc := append(aaa, bbb...)
	// ...를 통해 슬라이스의 각 요소 넘겨줄 수 있다.
	// 또 '...'를 가변인자 함수를 만들 때 사용 할 수 있다. (append도 이를활용해 구현됨)

	fmt.Println(ccc)

	/*
	 *	Reference
	 **/
	kkk := ccc
	kkk[0] = 999999
	fmt.Println(ccc, kkk)

	// vs. array (copy)
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 999
	fmt.Println(arr1, arr2)

	// want to copy slice
	yyy := []int{1, 2, 3, 4, 5}
	zzz := make([]int, 3)

	copy(zzz, yyy)

	fmt.Println(yyy, zzz)

	/*
	 *	partial slice --> reference
	 **/

	partial := yyy[0:2]
	fmt.Println("partial:", yyy, partial)
	fmt.Println(yyy[:])
	fmt.Println(yyy[0:])
	fmt.Println(yyy[:5])
	fmt.Println(yyy[0:len(yyy)])

	partial2 := yyy[0:2:cap(yyy)] // 세번째 숫자 --> slice 용량 지정(원본보다 클수 없음)
	fmt.Println("partial's capa:", cap(partial2))

}
