/*
 */
// Channels : CLOSING CHANNELS IN GO

// [CLOSING CHANNELS IN GO]
// Channels can be explicitly closed by a sender:
/*
	ch := make(chan int)

	// do some stuff with the channel
	close(ch)
*/

// [CHECKING IF A CHANNEL IS CLOSED]
// Similar to the ok value when accessing data in a map, receivers can check the ok
// value when receiving from a channel to test if a channel was closed.
/*
	v, ok := <-ch
	// ok is false if the channel is empty and closed.
*/

// [DON'T SEND ON A CLOSED CHANNEL]
// Sending on a closed channel will cause a panic. A panic on the main goroutine will
// cause the entire program to crash, and a panic in any other goroutine will cause
// that goroutine to crash.

// Closing isn't necessary. There's nothing wrong with leaving channels open, they'll
// still be garbage collected if they're unused. You should close channels to indicate
// explicitly to a receiver that nothing else is going to come across.

// ASSIGNMENT
// At Mailio we're all about keeping track of what our systems are up to with great
// logging and telemetry.

// The sendReports function sends out a batch of reports to our clients and reports
// back how many were sent across a channel. It closes the channel when it's done.

// Complete the countReports function. It should:

// - Use an infinite for loop to read from the channel:
// - If the channel is closed, break out of the loop
// - Otherwise, keep a running total of the number of reports sent
// - Return the total number of reports sent

package main

import (
	"fmt"
	"time"
)

func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

// don't touch below this line

func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func main() {
	test(3)
	test(4)
	test(5)
	test(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

// RESULTS:

// Start counting...
// Sent batch of 15 reports
// Sent batch of 38 reports
// Sent batch of 61 reports
// 114 reports sent!
// ========================
// Start counting...
// Sent batch of 15 reports
// Sent batch of 38 reports
// Sent batch of 61 reports
// Sent batch of 84 reports
// 198 reports sent!
// ========================
// Start counting...
// Sent batch of 15 reports
// Sent batch of 38 reports
// Sent batch of 61 reports
// Sent batch of 84 reports
// Sent batch of 107 reports
// 305 reports sent!
// ========================
// Start counting...
// Sent batch of 15 reports
// Sent batch of 38 reports
// Sent batch of 61 reports
// Sent batch of 84 reports
// Sent batch of 107 reports
// Sent batch of 130 reports
// 435 reports sent!
// ========================
