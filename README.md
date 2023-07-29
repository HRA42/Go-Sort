# Go Sort

## Overview

This Repo shows different sorting algorithms implemented in Golang.  
Every sorting algorithm is packed into its own folder and loaded into main. I also created a benchmark file. Don't run the entire file at once, it will fail due to the 11 minutes limit.

The main package does implement call for every package that in this module:

1. Array of, 100000 random integers between 0 and 9223372036854775807 (`math.MaxInt`)
2. Copy the created array, run Insertion Sort and measure the time it takes to sort the array
3. Copy the created array, run Bubble Sort and measure the time it takes to sort the array
4. Copy the created array, run Shell Sort and measure the time it takes to sort the array
5. Copy the created array, run Quick Sort and measure the time it takes to sort the array
6. Copy the created array, run Merge Sort and measure the time it takes to sort the array
7. Copy the created array, run the sorting algorithm build into go sort (from `sort pkg`) and measure the time it takes to sort the array#
8. Copy the created array, run Slice Sort (from `golang.org/x/exp/slices`) and measure the time it takes to sort the array

The benchmark file `./main_test.go` does the same, but run with the input cases:

1. 10
2. 100
3. 1000
4. 10000
5. 100000
6. 1000000
7. 10000000
8. 100000000

> [!WARNING]  
> For some sorting algorithms It's impossible to get a successful run (e.g. Bubble Sort)!

---

## Codesandbox

You can run this Repo in Codesandbox!  
  
[![Edit in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)]()
---

## Commands to run  

### run main.go

```bash
go run main.go
```

### run benchmark

```bash
go test -benchmem -run=^$ -bench ^Benchmark_createArray$
```
