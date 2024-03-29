package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Person -
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	// "application/xml"

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
