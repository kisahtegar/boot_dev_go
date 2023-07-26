// Loops : FIZZBUZZ

// ASSIGNMENT
// We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game,
// a coding exercise that has been dramatically overused in coding interviews across the world.

// Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line,
// but substitutes multiples of 3 for the text fizz and multiples of 5 for buzz. For multiples of 3 AND 5 print instead fizzbuzz.

package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

// RESULTS:

// 1
// 2
// fizz
// 4
// buzz
// fizz
// 7
// 8
// fizz
// buzz
// 11
// fizz
// 13
// 14
// fizzbuzz
// 16
// 17
// fizz
// 19
// buzz
// fizz
// 22
// 23
// fizz
// ....
