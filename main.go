/*
 */
// Pointer : POINTER RECEIVERS

// POINTER RECEIVERS
// A receiver type on a method can be a pointer.

// Methods with pointer receivers can modify the value to which the receiver points.
// Since methods often need to modify their receiver, pointer receivers are more
// common than value receivers.

// POINTER RECEIVER
/*
	type car struct {
		color string
	}

	func (c *car) setColor(color string) {
		c.color = color
	}

	func main() {
		c := car{
			color: "white",
		}
		c.setColor("blue")
		fmt.Println(c.color)
		// prints "blue"
	}
*/

// NON-POINTER RECEIVER
/*
	type car struct {
		color string
	}

	func (c car) setColor(color string) {
		c.color = color
	}

	func main() {
		c := car{
			color: "white",
		}
		c.setColor("blue")
		fmt.Println(c.color)
		// prints "white"
	}
*/
