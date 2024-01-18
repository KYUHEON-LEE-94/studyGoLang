package main

func main() {
	// 익명함수
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	print(result)

	//일급 함수
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

}

// 원형 정의
type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
