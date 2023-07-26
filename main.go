/*
 */
// Advanced Functions : DEFER

// DEFER
// The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically
// just before its enclosing function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until
// the surrounding function returns.

// Deferred functions are typically used to close database connections, file handlers and the like.

// For example:
/*
	// CopyFile copies a file from srcName to dstName on the local filesystem.
	func CopyFile(dstName, srcName string) (written int64, err error) {

		// Open the source file
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
		// Close the source file when the CopyFile function returns
		defer src.Close()

		// Create the destination file
		dst, err := os.Create(dstName)
		if err != nil {
			return
		}
		// Close the destination file when the CopyFile function returns
		defer dst.Close()

		return io.Copy(dst, src)
	}
*/
// In the above example, the src.Close() function is not called until after the CopyFile function was
// called but immediately before the CopyFile function returns.

// Defer is a great way to make sure that something happens at the end of a function, even if there
// are multiple return statements.

// ASSIGNMENT
// There is a bug in the logAndDelete function, fix it!

// This function should always delete the user from the user's map, which is a map that stores the
// user's name as keys. It also returns a log string that indicates to the caller some information
// about the user's deletion.

// To avoid bugs like this in the future, instead of calling delete before each return, just defer
// the delete once at the beginning of the function.

package main

import (
	"fmt"
	"sort"
)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		// delete(users, name)
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	// delete(users, name)
	return logDeleted
}

// don't touch below this line

type user struct {
	name   string
	number int
	admin  bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func main() {
	users := map[string]user{
		"john": {
			name:   "john",
			number: 18965554631,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 19875556452,
			admin:  true,
		},
		"breanna": {
			name:   "breanna",
			number: 98575554231,
			admin:  false,
		},
		"kade": {
			name:   "kade",
			number: 10765557221,
			admin:  false,
		},
	}

	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	test(users, "john")
	test(users, "santa")
	test(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")
}

// RESULTS:

// Initial users:
//  - breanna
//  - elon
//  - john
//  - kade
// ====================================
// Attempting to delete john...
// Log: admin deleted
// ====================================
// Attempting to delete santa...
// Log: user not found
// ====================================
// Attempting to delete kade...
// Log: user deleted
// ====================================
// Final users:
//  - breanna
//  - elon
