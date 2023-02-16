/*

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. 
The program should be written as a loop. Before entering the loop, the program should create an 
empty integer slice of size (length) 3. During each pass through the loop, the program prompts 
the user to enter an integer to be added to the slice. The program adds the integer to the slice, 
sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size 
to accommodate any number of integers which the user decides to enter. The program should only quit 
(exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/


package main

import (
	"fmt"
	"strings"
	"sort"
	"strconv"
)

func main() {
	var input string
	slice := make([]int, 0, 3) // create an empty integer slice of size (length) 3
	var counter int

	for {
		fmt.Print("Enter an integer or 'x' to exit: ")
		fmt.Scan(&input)

		if strings.ToLower(input) == "x" {
			break
		}

		if number, err := strconv.Atoi(input); err != nil {
			fmt.Println("Input string is not an integer number. Please try again.")
			continue
		} else {
			slice = append(slice, number)
			sort.Ints(slice) // sort the slice
			fmt.Println("Slice in the (increasing) order:", slice)
			counter++ // increasing the counter
		}
	}
}