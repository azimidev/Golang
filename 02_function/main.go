package main

import "fmt"

func main() {
	// var x float64
	// var y float64

	// x := 1.0
	// y := 2.0

	x, y := 1.0, 2.0

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	// var mean float64
	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T", mean, mean)
}
