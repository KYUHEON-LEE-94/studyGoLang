package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
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
