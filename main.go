package main

import (
	"fmt"

	"github.com/arispen/item-server/generator"
)

func main() {
	fmt.Println("Greetings, stranger.")
	item := generator.GenerateItem(1)

	fmt.Printf("%#v", item)
	fmt.Println("")

	for i := 0; ; i++ {
		item2 := generator.GenerateItem(1)
		if item2.Tier == generator.Unique {
			fmt.Printf("%#v", item2)
			fmt.Println(i)
			break
		}
	}
}
