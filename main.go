/*
 */
// Pointer : NIL POINTERS

// NIL POINTERS
// Pointers can be very dangerous.

// If a pointer points to nothing (the zero value of the pointer type) then dereferencing
// it will cause a runtime error (a panic) that crashes the program. Generally speaking,
// whenever you're dealing with pointers you should check if it's nil before trying to
// dereference it.

// ASSIGNMENT
// Let's make our profanity checker safe. Update the removeProfanity function. If message
// is nil, return early to avoid a panic. After all, there are no bad words to remove.

package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		if message == "" {
			removeProfanity(nil)
			fmt.Println("nil message detected")
		} else {
			removeProfanity(&message)
			fmt.Println(message)
		}
	}
}

func main() {
	messages := []string{
		"well shoot, this is awful",
		"",
		"dang robots",
		"dang them to heck",
		"",
	}

	messages2 := []string{
		"well shoot",
		"",
		"Allan is going straight to heck",
		"dang... that's a tough break",
		"",
	}

	test(messages)
	test(messages2)

}

// RESULTS:

// well *****, this is awful
// nil message detected
// **** robots
// **** them to ****
// nil message detected
// well *****
// nil message detected
// Allan is going straight to ****
// ****... that's a tough break
// nil message detected
