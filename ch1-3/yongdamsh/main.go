package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/O-L-K-L/go-study/ch1-3/yongdamsh/wc"
)

func main() {
	var s string

	// 이렇게 하면 한 단어 밖에 못 읽어들임
	// _, err := fmt.Scan("%s", &s)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s += scanner.Text()
		fmt.Println("Current word count:", wc.Calculate(s))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
