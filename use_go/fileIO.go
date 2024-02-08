package main

import (
	"io"
	"os"
)

func main() {
	filePath := "../../1.txt"
	//파일 생성
	fo, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	//파일 읽기
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buff := make([]byte, 1024)
	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		//끝이면 루프 종료
		if cnt == 0 {
			break
		}

		//쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

}
