package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was provided, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// if name was received, return a value that embeds the name
	// in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Returns a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
