// Slices : SLICES IN GO

// Slices always have an underlying array, though it isn't always specified explicitly.
// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]

// ASSIGNMENT
// Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

// Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input, and you've been provided with constants for the plan types at the top of the program.

// - If the plan is a pro plan, return all the strings from getMessageWithRetries().
// - If the plan is a free plan, return the first 2 strings from getMessageWithRetries().
// - If the plan isn't either of those, return an error that says unsupported plan.

package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// don't touch below this line

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("=====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func main() {
	test("Ozgur", 3, planFree)
	test("Jeff", 3, planPro)
	test("Sally", 2, planPro)
	test("Sally", 3, "no plan")
}

// RESULTS:

// sending to Ozgur...
// sending: "click here to sign up"
// sending: "pretty please click here"
// no response
// =====================================
// sending to Jeff...
// sending: "click here to sign up"
// sending: "pretty please click here"
// sending: "we beg you to sign up"
// no response
// =====================================
// sending to Sally...
// sending: "click here to sign up"
// sending: "pretty please click here"
// sending: "we beg you to sign up"
// they responded!
// =====================================
// sending to Sally...
// Error: unsupported plan
// =====================================
