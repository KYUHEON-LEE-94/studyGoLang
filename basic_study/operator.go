package main

import "fmt"

func main() {
	var k int = 10
	var p = &k
	fmt.Println(*p)
	fmt.Println(p)
}
