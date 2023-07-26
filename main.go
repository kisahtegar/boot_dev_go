/*
 */
// Pointer : INTRODUCTION TO POINTERS

// POINTERS
// Pointers hold the memory address of a value.
/*
	// The * syntax defines a pointer:
	var p *int
	A pointer's zero value is nil

	// The & operator generates a pointer to its operand.
	myString := "hello"
	myStringPtr = &myString

	// The * dereferences a pointer to gain access to the value
	fmt.Println(*myStringPtr) // read myString through the pointer
	*myStringPtr = "world"    // set myString through the pointer
*/
// Unlike C, Go has no pointer arithmetic

// [ JUST BECAUSE YOU CAN DOESN'T MEAN YOU SHOULD ]
//
// We're doing this exercise to understand that pointers can be used in this way.
// That said, pointers can be very dangerous. It's generally a better idea to have
// your functions accept non-pointers and return new values rather than mutating
// pointer inputs.

// [ ASSIGNMENT ]
// Complete the removeProfanity function.

// It should use the strings.ReplaceAll function to replace all instances of the
// following words in the input message with asterisks.
// - "dang" -> "****"
// - "shoot" -> "*****"
// - "heck" -> "****"

// It should mutate the value in the pointer and return nothing.

// Do not alter the function signature.

package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}

// RESULTS:

// well *****, this is awful
// **** robots
// **** them to ****
// well *****
// Allan is going straight to ****
// ****... that's a tough break
