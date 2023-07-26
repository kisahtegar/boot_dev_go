// Errors : THE ERROR PACKAGE

package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("No dividing by zero")
	}
	return x / y, nil
}

// don't edit below this line

func test(x, y float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}

// RESULTS:

// Dividing 10.00 by 0.00 ...
// No dividing by zero
// ====================================
// Dividing 10.00 by 2.00 ...
// Quotient: 5.00
// ====================================
// Dividing 15.00 by 30.00 ...
// Quotient: 0.50
// ====================================
// Dividing 6.00 by 3.00 ...
// Quotient: 2.00
// ====================================
