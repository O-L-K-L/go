package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := new(http.Client)
	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
	}
}
