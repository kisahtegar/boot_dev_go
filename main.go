/*
 */
// MAPS : MAPS

// MAPS
// Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

// The zero value of a map is nil.

// We can create a map by using a literal or by using the make() function:
/*
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites 24
	ages = map[string]int{
	"John": 37,
	"Mary": 21,
	}
*/

// The len() function works on a map, it returns the total number of key/value pairs.
/*
	ages = map[string]int{
	  "John": 37,
	  "Mary": 21,
	}
	fmt.Println(len(ages)) // 2
*/

// ASSIGNMENT
// We can speed up our contact-info lookups by using a map! Looking up a value in a map by its key is much faster
// than searching through a slice.

// Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of
// name -> user structs and potentially an error. A user struct just contains a user's name and phone number.

// If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

// The first name in the names slice matches the first phone number, and so on.

package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumbers := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers,
		}
	}
	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}

// RESULTS:

// Creating map...
// key: John, value:
//  - name: John
//  - number: 14355550987
// key: Bob, value:
//  - name: Bob
//  - number: 98765550987
// key: Jill, value:
//  - name: Jill
//  - number: 18265554567
// ====================================
// Creating map...
// invalid sizes
// ====================================
// Creating map...
// key: George, value:
//  - name: George
//  - number: 20955559812
// key: Sally, value:
//  - name: Sally
//  - number: 38385550982
// key: Rich, value:
//  - name: Rich
//  - number: 48265554567
// key: Sue, value:
//  - name: Sue
//  - number: 16045559873
// ====================================
