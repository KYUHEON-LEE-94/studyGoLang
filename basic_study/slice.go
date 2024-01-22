package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10

	// fmt 패키지의 Println 함수를 사용하여 슬라이스 전체를 출력
	fmt.Println(a)

	s := make([]int, 5, 10)
	println(len(s), cap(s)) //length 5, capacity 10

	var nilSlice []int
	if nilSlice == nil {
		println("Nil Slice")
	}
	println(len(nilSlice), cap(nilSlice)) // 모두 0

	var subSlice = []int{0, 1, 2, 3, 4, 5}
	subSlice = subSlice[2:5] // 2, 3, 4
	fmt.Println(subSlice)
	subSlice = subSlice[1:] // 3, 4
	fmt.Println(subSlice)

	appendSlice := []int{0, 1}
	appendSlice = append(appendSlice, 2)
	appendSlice = append(appendSlice, 3, 4, 5)
	fmt.Println(appendSlice)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)
	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력
}
