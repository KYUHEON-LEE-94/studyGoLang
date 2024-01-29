package main

//구조체
type Rect struct {
	width, height int
}

//메서드, Rect의 메서드로서 동작
func (r Rect) area() int {
	return r.height * r.width
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출
	area2 := rect.area2()
	println(area)
	println(area2)
}
