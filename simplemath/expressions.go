package simplemath

import (
	"errors"
)

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

// func Divide(p1, p2 float64) (float64, error) {
// 	if p2 == 0 {
// 		return math.NaN(), errors.New("cannot divide by zero")
// 	}
// 	return p1 / p2, nil
// }

// using named returns
func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}
	answer = p1 / p2
	return
}

// taking a bunch of parameters and dropping them into a slice
// needs to be the last parameter
func Sum(values ...float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}
