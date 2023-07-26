/*
 */
// Slices : RANGE

// RANGE
// Go provides syntactic sugar to iterate easily over elements of a slice:
/*
	for INDEX, ELEMENT := range SLICE {
	}
*/

// For example:
/*
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
	// 0 apple
	// 1 banana
	// 2 grape
*/
// ASSIGNMENT
// We need to be able to quickly detect bad words in the messages our system sends.
// Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return
// the index of the first bad word in the msg slice. This will help us filter out naughty words from our
// messaging system. If no bad words are found, return -1 instead.

// Use the range keyword.

package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

// don't touch below this line

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test(message, badWords)
}

// RESULTS:

// Scanning message: [hey there john] for bad words:
//  - crap
//  - shoot
//  - dang
//  - frick
// Index: -1
// ====================================
// Scanning message: [ugh oh my frick] for bad words:
//  - crap
//  - shoot
//  - dang
//  - frick
// Index: 3
// ====================================
// Scanning message: [what the shoot I hate that crap] for bad words:
//  - crap
//  - shoot
//  - dang
//  - frick
// Index: 2
// ====================================
