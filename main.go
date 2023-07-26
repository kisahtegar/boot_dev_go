/*
 */
// Advanced Functions : CLOSURES

// CLOSURES
// A closure is a function that references variables from outside its own function body.
// The function may access and assign to the referenced variables.

// In this example, the concatter() function returns a function that has reference to an
// enclosed doc value. Each successive call to harryPotterAggregator mutates that same
// doc variable.
/*
	func concatter() func(string) string {
		doc := ""
		return func(word string) string {
			doc += word + " "
			return doc
		}
	}
	func main() {
		harryPotterAggregator := concatter()
		harryPotterAggregator("Mr.")
		harryPotterAggregator("and")
		harryPotterAggregator("Mrs.")
		harryPotterAggregator("Dursley")
		harryPotterAggregator("of")
		harryPotterAggregator("number")
		harryPotterAggregator("four,")
		harryPotterAggregator("Privet")

		fmt.Println(harryPotterAggregator("Drive"))
		// Mr. and Mrs. Dursley of number four, Privet Drive
	}
*/

// ASSIGNMENT
// Keeping track of how many emails we send is mission-critical at Mailio. Complete the
// adder() function.

// It should return a function that adds its input (an int) to an enclosed sum value,
// then return the new sum. In other words, it keeps a running total of the sum variable
// within a closure.

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Println(bill)
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})
}

// RESULTS:

// You've sent 1 emails and it has cost you 45 cents
// You've sent 2 emails and it has cost you 77 cents
// You've sent 3 emails and it has cost you 120 cents
// You've sent 4 emails and it has cost you 132 cents
// You've sent 5 emails and it has cost you 166 cents
// You've sent 6 emails and it has cost you 220 cents
// ====================================
// You've sent 1 emails and it has cost you 12 cents
// You've sent 2 emails and it has cost you 24 cents
// You've sent 3 emails and it has cost you 1000 cents
// You've sent 4 emails and it has cost you 1012 cents
// You've sent 5 emails and it has cost you 1555 cents
// ====================================
// You've sent 1 emails and it has cost you 743 cents
// You've sent 2 emails and it has cost you 756 cents
// You've sent 3 emails and it has cost you 764 cents
// ====================================
