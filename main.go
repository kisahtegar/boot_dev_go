/*
 */
// Channels : CHANNELS

// [CHANNELS]
// Empty structs are often used as tokens in Go programs. In this context, a token
// is a unary value. In other words, we don't care what is passed through the channel.
// We care when and if it is passed.

// We can block and wait until something is sent on a channel using the following syntax
/*
	<-ch
*/
// This will block until it pops a single item off the channel, then continue, discarding
// the item.

// ASSIGNMENT
// Our Mailio server isn't able to boot up until it receives the signal that its databases
// are all online, and it learns about them being online by waiting for tokens (empty structs)
// on a channel.

// Complete the waitForDbs function. It should block until it receives numDBs tokens on the
// dbChan channel. Each time it reads a token the getDatabasesChannel goroutine will print
// a message to the console for you.

package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

// don't touch below this line

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

// RESULTS:

// Waiting for 3 databases...
// Database 1 is online
// Database 2 is online
// Database 3 is online
// All databases are online!
// =====================================
// Waiting for 4 databases...
// Database 1 is online
// Database 2 is online
// Database 3 is online
// Database 4 is online
// All databases are online!
// =====================================
// Waiting for 5 databases...
// Database 1 is online
// Database 2 is online
// Database 3 is online
// Database 4 is online
// Database 5 is online
// All databases are online!
// =====================================
