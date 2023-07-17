package main

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/hra42/go-sort/bubblesort"
	"github.com/hra42/go-sort/insertionsort"
	"github.com/hra42/go-sort/mergesort"
	"github.com/hra42/go-sort/quicksort"
	"github.com/hra42/go-sort/randomarray"
	"github.com/hra42/go-sort/shellsort"
	"golang.org/x/exp/slices"
)

func copy_arr(arr []int) []int {
	arr_copy := make([]int, len(arr))
	copy(arr_copy, arr)
	return arr_copy
}

func main() {
	// create Array
	start_creation := time.Now()
	arr := randomarray.Create_random_int_array(100000, math.MaxInt)
	duration_creation := time.Since(start_creation)
	fmt.Println("Creation: " + fmt.Sprint(duration_creation))

	// Insertion Sort
	insertion_arr := copy_arr(arr)
	start_insertionSort := time.Now()
	insertionsort.Sort(insertion_arr)
	duration_insertionSort := time.Since(start_insertionSort)
	fmt.Println("InsertionSort: " + fmt.Sprint(duration_insertionSort))

	// Bubble Sort
	bubble_arr := copy_arr(arr)
	start_bubbleSort := time.Now()
	bubblesort.Sort(bubble_arr)
	duration_bubbleSort := time.Since(start_bubbleSort)
	fmt.Println("BubbleSort: " + fmt.Sprint(duration_bubbleSort))

	// shell sort
	shell_arr := copy_arr(arr)
	start_shellSort := time.Now()
	shellsort.Sort(shell_arr)
	duration_shellSort := time.Since(start_shellSort)
	fmt.Println("ShellSort: " + fmt.Sprint(duration_shellSort))

	// quick sort
	quick_arr := copy_arr(arr)
	start_quickSort := time.Now()
	quicksort.Sort(quick_arr)
	duration_quickSort := time.Since(start_quickSort)
	fmt.Println("QuickSort: " + fmt.Sprint(duration_quickSort))

	// merge sort
	merge_arr := copy_arr(arr)
	start_mergeSort := time.Now()
	mergesort.Sort(merge_arr)
	duration_mergeSort := time.Since(start_mergeSort)
	fmt.Println("MergeSort: " + fmt.Sprint(duration_mergeSort))

	// go build in sort
	go_arr := copy_arr(arr)
	start_goSort := time.Now()
	sort.Ints(go_arr)
	duration_goSort := time.Since(start_goSort)
	fmt.Println("GoSort: " + fmt.Sprint(duration_goSort))

	slice_arr := copy_arr(arr)
	start_sliceSort := time.Now()
	slices.Sort(slice_arr)
	duration_sliceSort := time.Since(start_sliceSort)
	fmt.Println("SliceSort: " + fmt.Sprint(duration_sliceSort))
}
