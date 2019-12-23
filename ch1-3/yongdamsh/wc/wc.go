// Package wc 는 단어 개수를 세는 패키지
package wc

import (
	"github.com/webgenie/go-in-action/chapter3/words"
)

// Calculate 는 문자열을 입력 받아 단어의 개수를 세어 반환한다.
func Calculate(text string) int {
	return words.CountWords(text)
}
