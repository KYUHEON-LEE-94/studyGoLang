package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// GET 호출
	// resp, err := http.Get("https://www.naver.com/")
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// // 결과 출력
	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", string(data))

	// Request 객체 생성
	req2, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req2.Header.Add("User-Agent", "Crawler")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp2, err := client.Do(req2)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	// 결과 출력
	bytes, _ := io.ReadAll(resp2.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
}
