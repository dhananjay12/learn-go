/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputValues, _ := ReadValues()
	BubbleSort(inputValues)
	fmt.Println("Sorted values:")
	fmt.Println(inputValues)
}

// Bubble sort
func BubbleSort(inputValues []int) {
	for i := 0; i < len(inputValues); i++ {
		for j := 0; j < len(inputValues)-1; j++ {
			if inputValues[j] > inputValues[j+1] {
				Swap(inputValues, j)
			}
		}
	}
}

// Swap
func Swap(inputValues []int, i int) {
	temp := inputValues[i]
	inputValues[i] = inputValues[i+1]
	inputValues[i+1] = temp
}

func ReadValues() (inputValues []int, err error) {
	fmt.Println("Please input numbers(separated by space):")
	br := bufio.NewReader(os.Stdin)
	a, _, err := br.ReadLine()          // read a line
	ns := strings.Split(string(a), " ") // split the line by space
	// convert string to int
	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		inputValues = append(inputValues, n)
	}

	return inputValues, err
}
