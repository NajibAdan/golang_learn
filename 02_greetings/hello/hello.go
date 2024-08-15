package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set thwe properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// looop through 3 times to see the variations of the greetings returned
	// A slice of names.
	names := []string{"Najib", "Adan", "Teloo"}
	//Requet a greeting message
	messages, err := greetings.Hellos(names)

	//if an error was returned, print it to console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message to the console
	for _, name := range names {
		fmt.Println(messages[name])
	}
	fmt.Println(messages)
}
