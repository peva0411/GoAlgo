package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a slice containing pseudorandom numbers in [0, max].
func makeRandomSlice(numItems int, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomSlice := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		randomSlice[i] = random.Intn(max)
	}

	return randomSlice
}

// print at most numItems items.
func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func isSorted(slice []int) bool {
	isSorted := true

	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			isSorted = false
			break
		}
	}

	return isSorted
}

// verify that the slice is sorted.
func checkSorted(slice []int) {

	if isSorted(slice) {
		fmt.Println("The slice is sorted")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
}

// use bubble sort ot sort the slice
func bubbleSort(slice []int) {

	for {
		for i := 1; i < len(slice); i++ {
			if slice[i] < slice[i-1] {
				temp := slice[i-1]
				slice[i-1] = slice[i]
				slice[i] = temp
			}
		}

		if isSorted(slice) {
			break
		}
	}
}

func main() {

	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	//make and display the result
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	//sort and display the result
	bubbleSort(slice)
	printSlice(slice, 40)

	// verify that it's sorted
	checkSorted(slice)
}
