/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = 1/2 a t^2 + v0t + s0

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var a, v0, s0, t float64

	//Read the input values
	a = ReadValue("acceleration: a")
	v0 = ReadValue("initial velocity: v0")
	s0 = ReadValue("initial displacement: s0")
	t = ReadValue("time")

	//Generate the function
	fn := GenDisplaceFn(a, v0, s0)

	//Print the displacement after t seconds
	fmt.Printf("Displacement after %f seconds is %f", t, fn(t))

}

// Computes displacement as a function of time
// assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
// float64 argument which is the displacement travelled after time t.
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {

	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}

}

// Generic function returning the read value and an error.
func ReadValue(input_name string) (value float64) {

	var temp string
	var valid int

	read_input := bufio.NewScanner(os.Stdin)

	temp = input_name
	valid = 0
	for valid == 0 {
		fmt.Printf("\nPlease enter %s as a floating point: ", input_name)
		read_input.Scan()
		temp = read_input.Text()
		value, err := strconv.ParseFloat(temp, 64)
		if err == nil {
			valid = 1
			return value
		} else {
			fmt.Println("Invalid input. Floating point number required.")
		}
	}
	return value
}
