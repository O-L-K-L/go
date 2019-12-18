package main

import (
	"log"
	"os"

	_ "github.com/webgenie/go-in-action/chapter2/sample/matchers"
	"github.com/webgenie/go-in-action/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Holmes")
}
