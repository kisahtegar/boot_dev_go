/*
 */
// Channels : RANGE

// RANGE
// Similar to slices and maps, channels can be ranged over.
/*
	for item := range ch {
		// item is the next value received from the channel
	}
*/

// This example will receive values over the channel (blocking at each iteration if
// nothing new is there) and will exit only when the channel is closed.

// ASSIGNMENT
// It's that time again, Mailio is hiring and we've been assigned to do the interview.
// For some reason, the Fibonacci sequence is Mailio's interview problem of choice.
// We've been tasked with building a small toy program we can use in the interview.

// Complete the concurrrentFib function. It should:

// Create a new channel of ints
// Call fibonacci in a goroutine, passing it the channel and the number of Fibonacci
// numbers to generate, n
// Use a range loop to read from the channel and print out the numbers one by one,
// each on a new line

package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()
	for v := range chInts {
		fmt.Println(v)
	}
}

// don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch) // close channel
}

// RESULTS:

// Printing 10 numbers...
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// ==============================
// Printing 5 numbers...
// 0
// 1
// 1
// 2
// 3
// ==============================
// Printing 20 numbers...
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55
// 89
// 144
// 233
// 377
// 610
// 987
// 1597
// 2584
// 4181
// ==============================
// Printing 13 numbers...
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55
// 89
// 144
// ==============================
