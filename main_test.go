package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"time"

	"github.com/hra42/go-sort/bubblesort"
	"github.com/hra42/go-sort/insertionsort"
	"github.com/hra42/go-sort/mergesort"
	"github.com/hra42/go-sort/quicksort"
	"github.com/hra42/go-sort/randomarray"
	"github.com/hra42/go-sort/shellsort"
	"golang.org/x/exp/slices"
)

var table = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
	{input: 1000000},
	{input: 10000000},
	{input: 100000000},
}

func timeElapsed(input int, name string) func() {
	start := time.Now()
	return func() {
		timeElapsed := time.Since(start)
		fmt.Println("The function "+name+" took", timeElapsed, "for", input)
	}
}

func Benchmark_createArray(b *testing.B) {
	for _, v := range table {
		defer timeElapsed(v.input, "createArray")()
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			randomarray.Create_random_int_array(v.input, math.MaxInt)
		})
	}
}

func Benchmark_insertionSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "insertionSort")()
			insertionsort.Sort(arr)
		})
	}
}

func Benchmark_bubbleSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "bubbleSort")()
			bubblesort.Sort(arr)
		})
	}
}

func Benchmark_shellSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "shellSort")()
			shellsort.Sort(arr)
		})
	}
}

func Benchmark_quickSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "quickSort")()
			quicksort.Sort(arr)
		})
	}
}

func Benchmark_mergeSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "mergeSort")()
			mergesort.Sort(arr)
		})
	}
}

func Benchmark_goSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "goSort")()
			sort.Ints(arr)
		})
	}
}

func Benchmark_sliceSort(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			arr := randomarray.Create_random_int_array(v.input, math.MaxInt)
			b.ResetTimer()
			defer timeElapsed(v.input, "sliceSort")()
			slices.Sort(arr)
		})
	}
}
