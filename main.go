package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/tislais/go-functions/simplemath"
)

type MathExpr = string

const (
	AddExp      = MathExpr("banana")
	SubtractExp = MathExpr("subtract")
	MultiplyExp = MathExpr("multiply")
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

	sv := simplemath.NewSemanticVersion(666, 420, 69)

	// unnecessary because go does this for us
	// p := &sv
	// p.IncrementMinor()

	fmt.Println(simplemath.StringFunction(sv))

	// increment without pointer
	sv = sv.IncrementMajor()

	// increment with pointer
	sv.IncrementMinor()

	fmt.Println(sv.StringMethod())
	fmt.Println("----------")

	// get the reference to RoundTripCounter
	// without the &, we're missing method RoundTrip
	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	// anonymous function self-invoked
	func() {
		fmt.Println("My first anonymous function.")
	}()

	// anonymous function set to a variable
	a := func(name string) string {
		fmt.Printf("Hi, %s!\n", name)
		return name
	}
	a("Banana")

	addExp := mathExpression()
	fmt.Println(addExp(2, 3))

	addExp2 := mathExpression2()
	fmt.Println(addExp2(2, 2))

	addExp3 := mathExpression3(MultiplyExp)
	fmt.Println(addExp3(2, 1))

	fmt.Printf("%v\n", double(3, 2, mathExpression3(SubtractExp)))

	p2 := powerOfTwo()
	value := p2()
	println(value)

	for i := 0; i < 10; i++ {
		value = p2()
		println(value)
	}

	var funcs []func() int64

	// add 10 functions to the slice
	for i := 0; i < 10; i++ {
		// clean variable - capturing it in scope
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	// loop through the funcs we created
	// prints 100 - 10 times
	// because i keeps getting updated
	for _, f := range funcs {
		fmt.Printf("funcs: %v\n", f())
	}
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

// returning an anonymous function from a named function
func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

// returning a named function from a named function
func mathExpression2() func(float64, float64) float64 {
	return simplemath.Add
}

// return one of a variety of functions
func mathExpression3(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExp:
		return simplemath.Add
	case SubtractExp:
		return simplemath.Subtract
	case MultiplyExp:
		return simplemath.Multiply
	default:
		return func(f float64, f2 float64) float64 { return 0 }
	}
}

// takes a function as a parameter
func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

// maintains state
func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}
