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

	// Request 객체 생성
	req, err := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}

	//Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/json")

	// Client객체 생성
	client := &http.Client{}
	// Client객체에서 Request 실행
	resp, err := client.Do(req)
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
