/*
 */
// Advanced Functions : ANONYMOUS FUNCTIONS

// ANONYMOUS FUNCTIONS
// Anonymous functions are true to form in that they have no name. We've been using them
// throughout this chapter, but we haven't really talked about them yet.

// Anonymous functions are useful when defining a function that will only be used once
// or to create a quick closure.
/*
	// doMath accepts a function that converts one int into another
	// and a slice of ints. It returns a slice of ints that have been
	// converted by the passed in function.
	func doMath(f func(int) int, nums []int) []int {
		var results []int
		for _, n := range nums {
			results = append(results, f(n))
		}
		return results
	}

	func main() {
		nums := []int{1, 2, 3, 4, 5}

		// Here we define an anonymous function that doubles an int
		// and pass it to doMath
		allNumsDoubled := doMath(func(x int) int {
			return x + x
		}, nums)

		fmt.Println(allNumsDoubled)
		// prints:
		// [2 4 6 8 10]
	}
*/

// ASSIGNMENT
// Complete the printReports function.

// Call printCostReport once for each message. Pass in an anonymous function as the
// costCalculator that returns an int equal to twice the length of the input message.
// within a closure.

package main

import "fmt"

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(msg string) int {
			return len(msg) * 2
		}, message)
	}
}

// don't touch below this line

func test(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

// RESULTS:

// Message: "Here's Johnny!" Cost: 28 cents
// Message: "Go ahead, make my day" Cost: 42 cents
// Message: "You had me at hello" Cost: 38 cents
// Message: "There's no place like home" Cost: 52 cents
// ====================================
// Message: "Hello, my name is Inigo Montoya. You killed my father. Prepare to die." Cost: 140 cents
// Message: "May the Force be with you." Cost: 52 cents
// Message: "Show me the money!" Cost: 36 cents
// Message: "Go ahead, make my day." Cost: 44 cents
// ====================================
