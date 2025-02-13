/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t2 + vot + so

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


Test the program by running it twice and
testing it with two different sets of values for acceleration, initial velocity,
initial displacement, and time. Give 3 points if the program works correctly
for one set of values, and give 3 more points if the program works correctly
for the other set of values.

Examine the code. If the code contains a
function called GenDisplaceFn()
which takes three float64 arguments, acceleration a, initial velocity vo,
and initial displacement so and returns a function, then give
another 2 points. If the function returned by GenDisplaceFn() is used to compute the time, give another 2 points.
*/

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Printf("Enter the value for acceleration: ")
	_, err := fmt.Scan(&acceleration)

	if err != nil {
		fmt.Println("Invalid input. Start again and please enter a valid float.")
		os.Exit(0)
	}

	fmt.Printf("Enter the value for initial velocity: ")
	_, err2 := fmt.Scan(&initialVelocity)

	if err2 != nil {
		fmt.Println("Invalid input. Start again and please enter a valid float.")
		os.Exit(0)
	}

	fmt.Printf("Enter the value for initial displacement: ")
	_, err3 := fmt.Scan(&initialDisplacement)

	if err3 != nil {
		fmt.Println("Invalid input. Start again and please enter a valid float.")
		os.Exit(0)
	}

	fmt.Printf("Enter the value for time: ")
	_, err4 := fmt.Scan(&time)

	if err4 != nil {
		fmt.Println("Invalid input. Start again and please enter a valid float.")
		os.Exit(0)
	}

	fn := GenDisplaceFn(10, 2, 1)

	fmt.Printf("\nDisplacement is: ")

	fmt.Println(fn(time))
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
}
