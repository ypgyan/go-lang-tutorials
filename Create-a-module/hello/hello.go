package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Yan", "Rhana", "Nena"}
	greetings, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range greetings {
		fmt.Println(message)
	}
}
