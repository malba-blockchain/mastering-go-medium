/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers.

The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice.

You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.


Test the program by running it twice and testing it with a different sequence of integers each time. The first test sequence of integers should be all positive numbers and the second test should have at least one negative number. Give 3 points if the program works correctly for one test sequence, and give 3 more points if the program works correctly for the second test sequence.

Examine the code. If the code contains a function called BubbleSort() which takes a slice of integers as an argument, then give another 2 points.
If the code contains a function called Swap() function which takes two arguments, a slice of
integers and an index value i, then give another 2 points.
*/

package main

import (
	"fmt"
	"strconv"
)

// Write a Bubble Sort program in Go.
// The program should print the integers out on one line, in sorted order, from least to greatest.
func main() {
	var scannedInt string

	slice := make([]int, 0)

	fmt.Printf("Enter an integer to add to the sequence (up to 10), press 'x' to exit: ")
	for {

		fmt.Scan(&scannedInt)

		if scannedInt == "x" {
			break
		}

		intToAppend, err := strconv.Atoi(scannedInt)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		slice = append(slice, intToAppend)

		fmt.Println("Current ints in slice:")

		fmt.Println(slice)
		//The program should prompt the user to type in a sequence of up to 10 integers.
		if len(slice) >= 10 {
			break
		}
	}
	bubbleSort(slice)

	fmt.Println("\nSlice sorted using bubble:")
	fmt.Println(slice)
}

// you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

// You should write a Swap() function which performs this operation.
// Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
func swap(slice []int, j int) {
	slice[j], slice[j+1] = slice[j+1], slice[j]
}
