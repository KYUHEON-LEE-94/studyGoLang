package main

func main() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}

	if a == 11 {
		goto END
	}
	println(a)

END:
	println("End")

	i := 0

L1:
	for {

		if i == 0 {
			break L1
		}
	}

	println("OK")
}
