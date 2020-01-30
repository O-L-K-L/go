package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 물리적 코어를 여러개 사용하여 고루틴이 병렬적으로 실행되도록
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Go-routine")

	go func() {
		defer wg.Done() // 카운팅 세마포어를 이용해 실행이 완료되면 카운팅을 내림

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("on execute...")
	wg.Wait() // WaitGroup이 종료될 때까지 메인 함수가 기다림.

	fmt.Println("\nexit")
}
