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

func copyArr(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	return arrCopy
}

func main() {
	// create Array
	startCreation := time.Now()
	arr := randomarray.CreateRandomIntArray(100000, math.MaxInt)
	durationCreation := time.Since(startCreation)
	fmt.Println("Creation: " + fmt.Sprint(durationCreation))

	// Insertion Sort
	insertionArr := copyArr(arr)
	startInsertionSort := time.Now()
	insertionsort.Sort(insertionArr)
	durationInsertionSort := time.Since(startInsertionSort)
	fmt.Println("InsertionSort: " + fmt.Sprint(durationInsertionSort))

	// Bubble Sort
	bubbleArr := copyArr(arr)
	startBubbleSort := time.Now()
	bubblesort.Sort(bubbleArr)
	durationBubbleSort := time.Since(startBubbleSort)
	fmt.Println("BubbleSort: " + fmt.Sprint(durationBubbleSort))

	// Shell Sort
	shellArr := copyArr(arr)
	startShellSort := time.Now()
	shellsort.Sort(shellArr)
	durationShellSort := time.Since(startShellSort)
	fmt.Println("ShellSort: " + fmt.Sprint(durationShellSort))

	// quick sort
	quickArr := copyArr(arr)
	startQuickSort := time.Now()
	quicksort.Sort(quickArr)
	durationQuickSort := time.Since(startQuickSort)
	fmt.Println("QuickSort: " + fmt.Sprint(durationQuickSort))

	// merge sort
	mergeArr := copyArr(arr)
	startMergeSort := time.Now()
	mergesort.Sort(mergeArr)
	durationMergeSort := time.Since(startMergeSort)
	fmt.Println("MergeSort: " + fmt.Sprint(durationMergeSort))

	// go build in sort
	goArr := copyArr(arr)
	startGoSort := time.Now()
	sort.Ints(goArr)
	durationGoSort := time.Since(startGoSort)
	fmt.Println("GoSort: " + fmt.Sprint(durationGoSort))

	// slice sort
	sliceArr := copyArr(arr)
	startSliceSort := time.Now()
	slices.Sort(sliceArr)
	durationSliceSort := time.Since(startSliceSort)
	fmt.Println("SliceSort: " + fmt.Sprint(durationSliceSort))
}
