// Loops : LOOPS IN GO

package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}

// RESULTS:

// Sending 10 messages
// Bulk send complete! Cost = 10.45
// ===============================================================
// Sending 20 messages
// Bulk send complete! Cost = 21.90
// ===============================================================
// Sending 30 messages
// Bulk send complete! Cost = 34.35
// ===============================================================
// Sending 40 messages
// Bulk send complete! Cost = 47.80
// ===============================================================
// Sending 50 messages
// Bulk send complete! Cost = 62.25
// ===============================================================
