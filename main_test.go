package main

import (
	"fmt"
	"github.com/hra42/go-sort/bubblesort"
	"github.com/hra42/go-sort/insertionsort"
	"github.com/hra42/go-sort/mergesort"
	"github.com/hra42/go-sort/quicksort"
	"github.com/hra42/go-sort/randomarray"
	"github.com/hra42/go-sort/shellsort"
	"golang.org/x/exp/slices"
	"math"
	"sort"
	"testing"
	"time"
)

var testcases = []struct {
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

var tests = []struct {
	input    []int
	expected []int
}{
	{[]int{3, 2, 1, 5, 4}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 5, 5, 2, 2, 2}, []int{2, 2, 2, 5, 5, 5}},
}

func Benchmark_createArray(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_insertionSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				insertionsort.Sort(arr)
				duration += time.Since(start) // add duration of each loop iteration
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_bubbleSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				bubblesort.Sort(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_shellSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				shellsort.Sort(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_quickSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				quicksort.Sort(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_mergeSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				mergesort.Sort(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_goSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				sort.Ints(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func Benchmark_sliceSort(b *testing.B) {
	for _, v := range testcases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				slices.Sort(arr)
				duration += time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func TestCreateRandomIntArray(t *testing.T) {
	max := math.MaxInt

	for _, v := range testcases {
		res := randomarray.CreateRandomIntArray(v.input, max)
		if len(res) != v.input {
			t.Errorf("Expected length of %v, but got %v", v.input, len(res))
		}

		for _, v := range res {
			if v < 0 || v >= max {
				t.Errorf("Value %v should be within [0, %v)", v, max)
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range tests {
		result := insertionsort.Sort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Test fehlgeschlagen! Erwartet: %v, Erhalten: %v", test.expected, result)
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, test := range tests {
		result := bubblesort.Sort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Test fehlgeschlagen! Erwartet: %v, Erhalten: %v", test.expected, result)
			}
		}
	}
}

func TestShellSort(t *testing.T) {
	for _, test := range tests {
		result := shellsort.Sort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Test fehlgeschlagen! Erwartet: %v, Erhalten: %v", test.expected, result)
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		result := quicksort.Sort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Test fehlgeschlagen! Erwartet: %v, Erhalten: %v", test.expected, result)
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		result := mergesort.Sort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Test fehlgeschlagen! Erwartet: %v, Erhalten: %v", test.expected, result)
			}
		}
	}
}
