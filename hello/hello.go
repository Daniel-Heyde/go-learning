package main

import (
    "fmt"
	"log"
	"math/rand"

    "example.com/greetings"
)

func main() {
	// set properties of predefined logger, including
	// log entry prefix and a flag to disable printing
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Daniel",
		"Liz",
		"Fred",
	}

    // Get a greeting message and print it.
    // message, err := greetings.Hello(randomName())

    messages, err := greetings.Hellos(names)
	// if an error was returned, print to console and exit
	if err != nil {
		log.Fatal(err)
	}
	// otherwise, no errors
    fmt.Println(messages)
}

func randomName() string {
	names := []string{
		"Daniel",
		"Liz",
		"",
	}

	return names[rand.Intn(len(names))]
}