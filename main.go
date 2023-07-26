/*
 */
// Slices : VARIADIC

// VARIADIC
// Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

// A variadic function receives the variadic arguments as a slice.

/*
	func concat(strs ...string) string {
	    final := ""
	    // strs is just a slice of strings
	    for _, str := range strs {
	        final += str
	    }
	    return final
	}
*/

/*
	func main() {
	    final := concat("Hello ", "there ", "friend!")
	    fmt.Println(final)
	    // Output: Hello there friend!
	}
*/

// The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with
// space delimiters and a newline at the end.
/*
	func Println(a ...interface{}) (n int, err error)
*/

// SPREAD OPERATOR
// The spread operator allows us to pass a slice into a variadic function. The spread operator consists
// of three dots following the slice in the function call.
/*
	func printStrings(strings ...string) {
		for i := 0; i < len(strings); i++ {
			fmt.Println(strings[i])
		}
	}
	func main() {
	    names := []string{"bob", "sue", "alice"}
	    printStrings(names...)
	}
*/

// ASSIGNMENT
// We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.
// Complete the sum function so that returns the sum of all of its inputs.

// NOTE
// Make note of how the variadic inputs and the spread operator are used in the test suite.

package main

import "fmt"

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

// don't edit below this line

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}

// RESULTS:

// Summing 3 costs...
// Bill for the month: 6.00
// ===== END REPORT =====
// Summing 5 costs...
// Bill for the month: 15.00
// ===== END REPORT =====
// Summing 10 costs...
// Bill for the month: 55.00
// ===== END REPORT =====
// Summing 15 costs...
// Bill for the month: 120.00
// ===== END REPORT =====
