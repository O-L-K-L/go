package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("이메일을 전송합니다. %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // not user user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "kevin.koo",
			email: "kevin@line.com",
		},
		level: "low",
	}

	ad.user.notify()

	// 타입 임베딩
	ad.notify()
}
