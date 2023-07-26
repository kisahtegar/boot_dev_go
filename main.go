/*
 */
// Channels : CONCURRENCY

// CONCURRENCY

// WHAT IS CONCURRENCY?
// Concurrency is the ability to perform multiple tasks at the same time. Typically, our
// code is executed one line at a time, one after the other. This is called sequential
// execution or synchronous execution.

// [concurrency]
//
// If the computer we're running our code on has multiple cores, we can even execute
// multiple tasks at exactly the same time. If we're running on a single core, a single
// core executes code at almost the same time by switching between tasks very quickly.
// Either way, the code we write looks the same in Go and takes advantage of whatever
// resources are available.

// HOW DOES CONCURRENCY WORK IN GO?
//
// Go was designed to be concurrent, which is a trait fairly unique to Go. It excels
// at performing many tasks simultaneously safely using a simple syntax.

// There isn't a popular programming language in existence where spawning concurrent
// execution is quite as elegant, at least in my opinion.

// Concurrency is as simple as using the go keyword when calling a function:
/*
	go doSomething()
*/

// In the example above, doSomething() will be executed concurrently with the rest of
// the code in the function. The go keyword is used to spawn a new goroutine.

// [ASSIGNMENT]
//
// At Mailio we send a lot of network requests. Each email we send must go out over
// the internet. To serve our millions of customers, we need a single Go program to
// be capable of sending thousands of emails at once.

// Edit the sendEmail() function to execute its anonymous function concurrently so
// that the "received" message prints after the "sent" message.

package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")
}

// RESULTS:

// Email sent: 'Hello there Stacy!'
// Email received: 'Hello there Stacy!'
// ========================
// Email sent: 'Hi there John!'
// Email received: 'Hi there John!'
// ========================
// Email sent: 'Hey there Jane!'
// Email received: 'Hey there Jane!'
// ========================
