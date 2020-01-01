package main

import (
	"fmt"
)

type Todo map[string]string

func main() {
	todos := make([]Todo, 5)

	newTodo := Todo{"title": "Write in Go"}

	for index := range todos {
		todos[index] = newTodo
	}

	fmt.Println("list: ", todos)
	fmt.Println("cap: ", cap(todos))

	todos = append(todos, Todo{"title": "Being a Go Developer"})

	fmt.Println("list: ", todos)
	fmt.Println("cap: ", cap(todos))

	lastTodos := todos[len(todos)-1:]
	fmt.Println("list: ", lastTodos)
	fmt.Println("cap: ", cap(lastTodos))
}
