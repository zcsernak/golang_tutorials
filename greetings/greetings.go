package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name swas given, return and error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) // for testing
	return message, nil
}

// Hellos returns a map that associates eacho of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to assocaiate names with messages.
	messages := make(map[string]string)
	// Loop throug the received slice of names
	// calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrived message with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets inital values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying a rnaomd index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
