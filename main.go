/*
 */
// Slices : MAKE

// MAKE
// Most of the time we don't need to think about the underlying array of a slice.
// We can create a new slice using the make function:
/*
	// func make([]T, len, cap) []T
	mySlice := make([]int, 5, 10)

	// the capacity argument is usually omitted and defaults to the length
	mySlice := make([]int, 5)
	Slices created with make will be filled with the zero value of the type.

*/

// If we want to create a slice with a specific set of values, we can use a slice literal:
/*
	mySlice := []string{"I", "love", "go"}
	Note that the array brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.
*/

// LENGTH
// The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:
/*
	mySlice := []string{"I", "love", "go"}
	fmt.Println(len(mySlice)) // 3
*/

// CAPACITY
// The capacity of a slice is the number of elements in the underlying array, counting from the first element
// in the slice. It is accessed using the built-in cap() function:
/*
	mySlice := []string{"I", "love", "go"}
	fmt.Println(cap(mySlice)) // 3
*/

// Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry
// about the capacity of a slice because it will automatically grow as needed.

// ASSIGNMENT
// We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

// If we know the rough size of a slice before we fill it up, we can make our program faster by creating the
// slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new
// underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later,
// but it means we can avoid all the expensive resizing we know that we'll need.

// Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

// Preallocate a slice for the message costs of the same length as the messages slice.
// Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the
// message in the messages slice at the same index. The cost of a message is the length of the message
// multiplied by 0.01.

package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

// don't edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}

// RESULTS:

// Messages:
//  - Welcome to the movies!
//  - Enjoy your popcorn!
//  - Please don't talk during the movie!
// Costs:
//  - 0.22
//  - 0.19
//  - 0.35
// ===== END REPORT =====
// Messages:
//  - I don't want to be here anymore
//  - Can we go home?
//  - I'm hungry
//  - I'm bored
// Costs:
//  - 0.31
//  - 0.15
//  - 0.10
//  - 0.09
// ===== END REPORT =====
// Messages:
//  - Hello
//  - Hi
//  - Hey
//  - Hi there
//  - Hey there
//  - Hi there
//  - Hello there
//  - Hey there
//  - Hello there
//  - General Kenobi
// Costs:
//  - 0.05
//  - 0.02
//  - 0.03
//  - 0.03
//  - 0.08
//  - 0.09
//  - 0.08
//  - 0.11
//  - 0.09
//  - 0.11
//  - 0.14
// ===== END REPORT =====
