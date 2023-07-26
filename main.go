// Loops : OMITTING CONDITIONS FROM A FOR LOOP IN GO

// ASSIGNMENT
// Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.
//
// Each message costs 1.0, plus an additional fee. The fee structure is:
//
// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03
//
// BROWSER FREEZE
// If you lock up your browser by creating an infinite loop that isn't breaking, just click the cancel button.

package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.1 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}

// RESULTS:

// Threshold: 10.00
// Maximum messages that can be sent: = 7
// ===============================================================
// Threshold: 20.00
// Maximum messages that can be sent: = 12
// ===============================================================
// Threshold: 30.00
// Maximum messages that can be sent: = 16
// ===============================================================
// Threshold: 40.00
// Maximum messages that can be sent: = 20
// ===============================================================
// Threshold: 50.00
// Maximum messages that can be sent: = 23
// ===============================================================
