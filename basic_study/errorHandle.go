package main

import (
	"log"
	"os"
)

func otherFunc() {
	return 2
}

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	_, err2 := otherFunc() //임의의 함수 호출

	switch err2.(type) {
	default: // no error
		println("ok")
	case MyError:
		log.Print("Log my error")
	case error:
		log.Fatal(err.Error())
	}
}
