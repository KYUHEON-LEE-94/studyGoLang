package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	//Add()안의 횟수만큼 goroutine 함수가 호출되어야함.
	wg.Add(999)

	counter := 0

	// goroutine 생성 및 다음 익명 함수의 작업을 할당
	for i := 0; i < 999; i++ {
		go func() {
			wg.Done()
			counter++
		}()
	}

	wg.Wait() // goroutine 작업의 완료까지 대기
	fmt.Println("counter :", counter)

}
