/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the
array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Create a wait group to wait for all goroutines to finish
var wg sync.WaitGroup

func main() {
	inputValues, _ := ReadValues()

	// Partition the array into 4 parts
	partitionSize := len(inputValues) / 4
	partition1 := inputValues[0:partitionSize]
	partition2 := inputValues[partitionSize : partitionSize*2]
	partition3 := inputValues[partitionSize*2 : partitionSize*3]
	partition4 := inputValues[partitionSize*3 : len(inputValues)]

	// Sort each partition
	wg.Add(4)
	go Sort(partition1)
	go Sort(partition2)
	go Sort(partition3)
	go Sort(partition4)

	// Wait for all goroutines to finish
	wg.Wait()
	// Merge the 4 sorted subarrays into one large sorted array
	outputValues := append(partition1, partition2...)
	outputValues = append(outputValues, partition3...)
	outputValues = append(outputValues, partition4...)
	sort.Ints(outputValues)

	fmt.Println("Sorted values:")
	fmt.Println(outputValues)
}

func Sort(unsortedSLices []int) []int {
	sort.Ints(unsortedSLices)
	wg.Done()
	return unsortedSLices
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
