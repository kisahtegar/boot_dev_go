/*
 */
// Channels : CHANNELS

// [CHANNELS]
// Channels are a typed, thread-safe queue. Channels allow different goroutines to
// communicate with each other.

// [CREATE A CHANNEL]
// Like maps and slices, channels must be created before use. They also use the same
// make keyword:
/*
	ch := make(chan int)

*/

// [SEND DATA TO A CHANNEL]
/*
	ch <- 69
*/
// The <- operator is called the channel operator. Data flows in the direction of the
// arrow. This operation will block until another goroutine is ready to receive the value.

// RECEIVE DATA FROM A CHANNEL
/*
	v := <-ch
*/
// This reads and removes a value from the channel and saves it into the variable v.
// This operation will block until there is a value in the channel to be read.

// [BLOCKING AND DEADLOCKS]
// A deadlock is when a group of goroutines are all blocking so none of them can continue.
// This is a common bug that you need to watch out for in concurrent programming.

// [ASSIGNMENT]
// Run the program. You'll see that it deadlocks and never exits. The filterOldEmails
// function is trying to send on a channel, but no other goroutines are running that
// can accept the value from the channel.

// Fix the deadlock by spawning a goroutine to send the "is old" values.

package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		sendIsOld(isOldChan, emails)
	}()

	// because the readers doesnt happend yet, and deadlock.
	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

// TEST SUITE -- Don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails []email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

type email struct {
	body string
	date time.Time
}

func test(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

func main() {
	test([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Today is the day!",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Letsa Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
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
