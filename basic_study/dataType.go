package main

import "fmt"

func main() {
	//복수 라인.
	rawLiteral := `아리랑 \n
	아리랑 \n
	아라리요`

	//이중인용부호
	interLiteral := "아리랑아리랑\n아라리요"
	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes, str2)
}
