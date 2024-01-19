package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}
	println(a1[0], a2[0])

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10
	println(multiArray[0][1][2])

	var ma = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	println(ma[1][2])
}
