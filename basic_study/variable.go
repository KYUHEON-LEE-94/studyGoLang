package main

import "fmt"

func main() {
	var a int
	var f float = 11.2
	a = 3
	var c string
	var i, j, k int = 1, 2, 3
	var i = 1 //추론도 가능
	var s = "Hi"
	fmt.Println(i, j, k)

	const c int = 10
	const s string = "Hi"
	fmt.Println(c, s)
}
