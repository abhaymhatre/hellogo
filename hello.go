package main

import (
	"fmt"
	"log"

	"github.com/abhaymhatre/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Abhay")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Abhay", "Anagha", "Arisha"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
