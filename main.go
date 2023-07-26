/*
 */
// Advanced Functions : CURRYING

// CURRYING
// Function currying is the practice of writing a function that takes a function (or functions) as input,
// and returns a new function.

// For example:
/*
	func main() {
		squareFunc := selfMath(multiply)
		doubleFunc := selfMath(add)

		fmt.Println(squareFunc(5))
		// prints 25

		fmt.Println(doubleFunc(5))
		// prints 10
	}

	func multiply(x, y int) int {
		return x * y
	}

	func add(x, y int) int {
		return x + y
	}

	func selfMath(mathFunc func(int, int) int) func (int) int {
		return func(x int) int {
			return mathFunc(x, x)
		}
	}
*/
// In the example above, the selfMath function takes in a function as its parameter, and returns a function
// that itself returns the value of running that input function on its parameter.

// ASSIGNMENT
// The Mailio API needs a very robust error-logging system so we can see when things are going awry in the
// back-end system. We need a function that can create a custom "logger" (a function that prints to the console)
// given a specific formatter.

// Complete the getLogger function. It should return a new function that prints the formatted inputs using the
// given formatter function. The inputs should be passed into the formatter function in the order they are given
// to the logger function.

package main

import (
	"errors"
	"fmt"
)

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}

// don't touch below this line

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}

// RESULTS:

// Logs:
// Error on database server: out of memory
// Error on database server: cpu is pegged
// Error on database server: networking issue
// Error on database server: invalid syntax
// ====================================
// Logs:
// Error on mail server, email too large
// Error on mail server, non alphanumeric symbols found
// ====================================
