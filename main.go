package main

import (
	"fmt"

	"github.com/tislais/go-functions/simplemath"
)

func main() {
	fmt.Printf("add: %f\n", simplemath.Add(6, 2))
	fmt.Printf("subtract: %f\n", simplemath.Subtract(6, 2))
	fmt.Printf("multiply: %f\n", simplemath.Multiply(6, 2))

	answer, err := simplemath.Divide(6, 0)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("divide: %f\n", answer)
	}

	numbers := []float64{380, 40, 0.666}

	// spreading a slice in as arguments
	total := simplemath.Sum(numbers...)
	fmt.Printf("%f\n", total)
}
