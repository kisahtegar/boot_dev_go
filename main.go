/*
 */
// Pointer : POINTER RECEIVER CODE

// POINTER RECEIVER CODE
// Methods with pointer receivers don't require that a pointer is used to call the
// method. The pointer will automatically be derived from the value.
/*
	type circle struct {
		x int
		y int
		radius int
	}

	func (c *circle) grow(){
		c.radius *= 2
	}

	func main(){
		c := circle{
			x: 1,
			y: 2,
			radius: 4,
		}

		// notice c is not a pointer in the calling function
		// but the method still gains access to a pointer to c
		c.grow()
		fmt.Println(c.radius)
		// prints 8
	}
*/

// ASSIGNMENT
// Fix the bug in the code so that setMessage sets the message field of the given
// email structure, and the new value persists outside the scope of the setMessage
// method.

package main

import (
	"fmt"
)

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage("this is my second draft")
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")

}

// RESULTS:

// -- before --
// message: this is my first draft
// fromAddress: sandra@mailio-test.com
// toAddress: bullock@mailio-test.com
// -- end before --
// -- after --
// message: this is my second draft
// fromAddress: sandra@mailio-test.com
// toAddress: bullock@mailio-test.com
// -- end after --
// ==========================
// -- before --
// message: this is my third draft
// fromAddress: sandra@mailio-test.com
// toAddress: bullock@mailio-test.com
// -- end before --
// -- after --
// message: this is my second draft
// fromAddress: sandra@mailio-test.com
// toAddress: bullock@mailio-test.com
// -- end after --
// ==========================
