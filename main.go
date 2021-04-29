package main

import (
	"fmt"
)

func main() {
	fmt.Println("Greetings, stranger.")
	item := GenerateItem()
	fmt.Printf("%#v", item)
	fmt.Println("")
}
