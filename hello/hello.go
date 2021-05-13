package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main () {
	log.SetFlags(0)
	log.SetPrefix("Greetings: ")

	names := []string{"yuvraj", "kabira", "lala"}
	message, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}