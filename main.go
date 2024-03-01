package main

import (
	"fmt"
	"log"

	"nosuparser/parser"
)

func main() {

	a, err := parser.ParseNosuNews(1, 10)

	if err != nil {
		log.Println(a)
	}

	if len(a) != 0 {
		for _, news := range a {
			fmt.Printf("%s\n", news.Title)
		}
	}

}
