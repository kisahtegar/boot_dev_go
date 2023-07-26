/*
 */
// Slices : SLICE OF SLICES

// SLICE OF SLICES
// Slices can hold other slices, effectively creating a matrix, or a 2D slice.
/*
	rows := [][]int{}
*/

// ASSIGNMENT
// We support various graphs and dashboards on Textio that display message analytics to our users.
// The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

// Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice
// of integers where the value of each cell is i * j where i and j are the indexes of the row and column respectively.

// For example, a 10x10 matrix would look like this:
// [0 0 0 0 0 0 0 0 0 0]
// [0 1 2 3 4 5 6 7 8 9]
// [0 2 4 6 8 10 12 14 16 18]
// [0 3 6 9 12 15 18 21 24 27]
// [0 4 8 12 16 20 24 28 32 36]
// [0 5 10 15 20 25 30 35 40 45]
// [0 6 12 18 24 30 36 42 48 54]
// [0 7 14 21 28 35 42 49 56 63]
// [0 8 16 24 32 40 48 56 64 72]
// [0 9 18 27 36 45 54 63 72 81]

package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}

// RESULTS:

// Creating 3 x 3 matrix...
// [0 0 0]
// [0 1 2]
// [0 2 4]
// ===== END REPORT =====
// Creating 5 x 5 matrix...
// [0 0 0 0 0]
// [0 1 2 3 4]
// [0 2 4 6 8]
// [0 3 6 9 12]
// [0 4 8 12 16]
// ===== END REPORT =====
// Creating 10 x 10 matrix...
// [0 0 0 0 0 0 0 0 0 0]
// [0 1 2 3 4 5 6 7 8 9]
// [0 2 4 6 8 10 12 14 16 18]
// [0 3 6 9 12 15 18 21 24 27]
// [0 4 8 12 16 20 24 28 32 36]
// [0 5 10 15 20 25 30 35 40 45]
// [0 6 12 18 24 30 36 42 48 54]
// [0 7 14 21 28 35 42 49 56 63]
// [0 8 16 24 32 40 48 56 64 72]
// [0 9 18 27 36 45 54 63 72 81]
// ===== END REPORT =====
// Creating 15 x 15 matrix...
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
// [0 2 4 6 8 10 12 14 16 18 20 22 24 26 28]
// [0 3 6 9 12 15 18 21 24 27 30 33 36 39 42]
// [0 4 8 12 16 20 24 28 32 36 40 44 48 52 56]
// [0 5 10 15 20 25 30 35 40 45 50 55 60 65 70]
// [0 6 12 18 24 30 36 42 48 54 60 66 72 78 84]
// [0 7 14 21 28 35 42 49 56 63 70 77 84 91 98]
// [0 8 16 24 32 40 48 56 64 72 80 88 96 104 112]
// [0 9 18 27 36 45 54 63 72 81 90 99 108 117 126]
// [0 10 20 30 40 50 60 70 80 90 100 110 120 130 140]
// [0 11 22 33 44 55 66 77 88 99 110 121 132 143 154]
// [0 12 24 36 48 60 72 84 96 108 120 132 144 156 168]
// [0 13 26 39 52 65 78 91 104 117 130 143 156 169 182]
// [0 14 28 42 56 70 84 98 112 126 140 154 168 182 196]
// ===== END REPORT =====
