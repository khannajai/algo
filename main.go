package main

import (
	"fmt"
	"github.com/khannajai/algo/sort"
)

func main() {
	mylist := []int{5, 4, 3, 2, 1}

	// mergesort
	sortedlist := sort.Mergesort(mylist)
	fmt.Println(sortedlist)

	// quicksort
	sortedlist = sort.Quicksort(mylist)
	fmt.Println(sortedlist)

	// selectionsort
	sortedlist = sort.Selectionsort(mylist)
	fmt.Println(sortedlist)
}
