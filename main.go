package main

import (
	"fmt"

	"github.com/arispen/item-server/generator"
)

func main() {
	fmt.Println("Greetings, stranger.")
	item := generator.GenerateItem()
	item2 := generator.GenerateItem()
	fmt.Printf("%#v", item)
	fmt.Println("")
	fmt.Printf("%#v", item2)
	fmt.Println("")
}
