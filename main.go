/*
 */
// Slices : APPEND

// APPEND
// The built-in append function is used to dynamically add elements to a slice:

/*
	func append(slice []Type, elems ...Type) []Type
	If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.
*/

// Notice that append() is variadic, the following are all valid:
/*
	slice = append(slice, oneThing)
	slice = append(slice, firstThing, secondThing)
	slice = append(slice, anotherSlice...)
*/

// ASSIGNMENT
// We've been asked to "bucket" costs for an entire month into the cost occurring on each day of the month.

// Complete the getCostsByDay function. It should return a slice of float64s, where each element is the total cost
// for that day. The length of the slice should be equal to the number of days represented in the costs slice,
// including any days that have no costs, up to the last day represented in the slice.

// Days are numbered and start at 0.

// EXAMPLE
// Input in day, cost format:
/*
	[]cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}
*/
// I would expect this result:
/*
	[]float64{
		4.0,
		5.2,
		0.0,
		0.0,
		0.0,
		2.5,
	}
*/

package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
			// fmt.Println(costsByDay)
		}
		costsByDay[cost.day] += cost.value
		fmt.Println(costsByDay)
	}
	return costsByDay
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}

// RESULTS:

// Creating daily buckets for 7 costs...
// Costs by day:
//  - Day 0: 1.00
//  - Day 1: 5.10
//  - Day 2: 2.50
//  - Day 3: 6.30
//  - Day 4: 3.34
// ===== END REPORT =====
// Creating daily buckets for 11 costs...
// Costs by day:
//  - Day 0: 1.00
//  - Day 1: 3.60
//  - Day 2: 5.20
//  - Day 3: 3.10
//  - Day 4: 56.34
//  - Day 5: 0.00
//  - Day 6: 0.00
//  - Day 7: 0.00
//  - Day 8: 0.00
//  - Day 9: 0.00
//  - Day 10: 2.00
//  - Day 11: 0.00
//  - Day 12: 0.00
//  - Day 13: 2.34
//  - Day 14: 0.00
//  - Day 15: 0.00
//  - Day 16: 0.00
//  - Day 17: 0.00
//  - Day 18: 0.00
//  - Day 19: 0.00
//  - Day 20: 0.00
//  - Day 21: 0.00
//  - Day 22: 0.00
//  - Day 23: 0.00
//  - Day 24: 0.00
//  - Day 25: 2.34
//  - Day 26: 0.00
//  - Day 27: 0.00
//  - Day 28: 1.34
//  - Day 29: 0.00
//  - Day 30: 4.34
// ===== END REPORT =====
