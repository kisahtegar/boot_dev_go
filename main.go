/*
 */
// MAPS : MUTATIONS

// MUTATIONS
/*
	// INSERT AN ELEMENT
	m[key] = elem
	// GET AN ELEMENT
	elem = m[key]
	// DELETE AN ELEMENT
	delete(m, key)
	// CHECK IF A KEY EXISTS
	elem, ok := m[key]
*/
// If key is in m, then ok is true. If not, ok is false.
// If key is not in the map, then elem is the zero value for the map's element type.

// ASSIGNMENT
// It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

// Complete the deleteIfNecessary function.

// If the user doesn't exist in the map, return the error not found.
// If they exist but aren't scheduled for deletion, return deleted as false with no errors.
// If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
// NOTE ON PASSING MAPS
// Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write,
// we can make changes to the original, we don't have a copy.

package main

import (
	"errors"
	"fmt"
	"sort"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

// don't touch below this line

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func main() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	test(users, "john")
	test(users, "musk")
	test(users, "santa")
	test(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}

// RESULTS:

// Attempting to delete john...
// Deleted: john
// ====================================
// Attempting to delete musk...
// not found
// ====================================
// Attempting to delete santa...
// not found
// ====================================
// Attempting to delete kade...
// Did not delete: kade
// ====================================
// Final map keys:
//  -  breanna
//  -  elon
//  -  kade
