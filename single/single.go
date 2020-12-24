package single

import (
	"fmt"
	"math/rand"
	"time"
)

// Error is error of single moduele
type Error struct {
	when time.Time
	what string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s Error!!! at %v", e.what, e.when)
}

// Hello returns a greeting for the named person.
func Hello(name string, short bool) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, Error{time.Now(), "empty name"}
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(short), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat(short bool) string {
	// A slice of message formats.
	shortFormats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v! ",
		"Hail, %v! Well met!",
	}

	longFormats := []string{
		"Hi, %v. Welcome! What a lovely day.",
		"Great to see you, %v! You look great",
		"Hi, %v! How are you doing?",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	if short {
		return shortFormats[rand.Intn(len(shortFormats))]
	}
	return longFormats[rand.Intn(len(longFormats))]
}
