package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := int64(10)

	for i := 0; i < 10; i++ {
		go func() {
			cnt--
		}()
	}

	if cnt > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(cnt)
}
