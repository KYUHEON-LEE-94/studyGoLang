package main

func main() {
	total := sum(1, 7, 8, 5, 9)
	println(total)

	count, total2 := sum2(1, 7, 8, 5, 9)
	println(count, total2)

	count2, total3 := sum3(1, 7, 8, 5, 9)
	println(count2, total3)
}

func sum(nums ...int) int {
	s := 0
	for _, number := range nums {
		s += number
	}
	return s
}

func sum2(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
