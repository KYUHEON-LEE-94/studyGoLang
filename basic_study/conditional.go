package main

import "fmt"

func main() {
	k := 1
	if k == 1 {
		fmt.Println(k)
	}
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	var test string
	checkType(test)

}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

// interface{}는 일종의 제너릭
func checkType(value interface{}) {
	// 값의 타입을 확인
	switch v := value.(type) {
	case int:
		fmt.Println("정수 타입:", v)
	case float64:
		fmt.Println("부동소수점 타입:", v)
	case string:
		fmt.Println("문자열 타입:", v)
	default:
		fmt.Println("알 수 없는 타입")
	}
}
