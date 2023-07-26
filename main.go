/*
 */
// Channels : BUFFERED CHANNELS

// [BUFFERED CHANNELS]
// Channels can optionally be buffered.

// [CREATING A CHANNEL WITH A BUFFER]
// You can provide a buffer length as the second argument to make() to create a
// buffered channel:

// ch := make(chan int, 100)
// Sending on a buffered channel only blocks when the buffer is full.

// Receiving blocks only when the buffer is empty.

// ASSIGNMENT
// We want to be able to send emails in batches. A writing goroutine will write an
// entire batch of email messages to a buffered channel, and later, once the channel
// is full, a reading goroutine will read all of the messages from the channel and
// send them out to our clients.

// Complete the addEmailsToQueue function. It should create a buffered channel with a
// buffer large enough to store all of the emails it's given. It should then write the
// emails to the channel in order, and finally return the channel.

package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}

// RESULTS:

// Adding 2 emails to queue...
// Sending emails...
// Sending email: Hello John, tell Kathy I said hi
// Sending email: Whazzup bruther
// ==========================================
// Adding 3 emails to queue...
// Sending emails...
// Sending email: I find that hard to believe.
// Sending email: When? I don't know if I can
// Sending email: What time are you thinking?
// ==========================================
// Adding 4 emails to queue...
// Sending emails...
// Sending email: She says hi!
// Sending email: Yeah its tomorrow. So we're good.
// Sending email: Cool see you then!
// Sending email: Bye!
// ==========================================
