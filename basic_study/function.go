package main

func main() {
	msg := "Hello"
	say(&msg)
	println(msg)

	say2("This", "is", "a", "book")
	say2("Hi")
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}

func say2(msg ...string) {
	for _, s := range msg {
		println(s)
	}

}
