package main

import (
	"fmt"
	"github.com/khannajai/algo/sort"
)

func main() {
	mylist := []int{5, 4, 3, 2, 1}

	// mergesort
	sortedlist := sort.MergeSort(mylist)
	fmt.Println(sortedlist)

	// quicksort
	sortedlist = sort.QuickSort(mylist)
	fmt.Println(sortedlist)

	// selectionsort
	sortedlist = sort.SelectionSort(mylist)
	fmt.Println(sortedlist)

	// heapsort
	sortedlist = sort.HeapSort(mylist)
	fmt.Println(sortedlist)
}
