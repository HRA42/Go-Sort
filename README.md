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
  
[![Edit in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)](https://codesandbox.io/p/github/HRA42/go-sort/main?layout=%257B%2522sidebarPanel%2522%253A%2522EXPLORER%2522%252C%2522rootPanelGroup%2522%253A%257B%2522direction%2522%253A%2522horizontal%2522%252C%2522contentType%2522%253A%2522UNKNOWN%2522%252C%2522type%2522%253A%2522PANEL_GROUP%2522%252C%2522id%2522%253A%2522ROOT_LAYOUT%2522%252C%2522panels%2522%253A%255B%257B%2522type%2522%253A%2522PANEL_GROUP%2522%252C%2522contentType%2522%253A%2522UNKNOWN%2522%252C%2522direction%2522%253A%2522vertical%2522%252C%2522id%2522%253A%2522clkolkcbj00bc356v0yvb20q4%2522%252C%2522sizes%2522%253A%255B70%252C30%255D%252C%2522panels%2522%253A%255B%257B%2522type%2522%253A%2522PANEL_GROUP%2522%252C%2522contentType%2522%253A%2522EDITOR%2522%252C%2522direction%2522%253A%2522horizontal%2522%252C%2522id%2522%253A%2522EDITOR%2522%252C%2522panels%2522%253A%255B%257B%2522type%2522%253A%2522PANEL%2522%252C%2522contentType%2522%253A%2522EDITOR%2522%252C%2522id%2522%253A%2522clkolkcbj00b7356v5kjcvj7e%2522%257D%255D%252C%2522sizes%2522%253A%255B100%255D%257D%252C%257B%2522type%2522%253A%2522PANEL_GROUP%2522%252C%2522contentType%2522%253A%2522SHELLS%2522%252C%2522direction%2522%253A%2522horizontal%2522%252C%2522id%2522%253A%2522SHELLS%2522%252C%2522panels%2522%253A%255B%257B%2522type%2522%253A%2522PANEL%2522%252C%2522contentType%2522%253A%2522SHELLS%2522%252C%2522id%2522%253A%2522clkolkcbj00bb356v2r1kwyws%2522%257D%255D%252C%2522sizes%2522%253A%255B100%255D%257D%255D%257D%252C%257B%2522type%2522%253A%2522PANEL_GROUP%2522%252C%2522contentType%2522%253A%2522DEVTOOLS%2522%252C%2522direction%2522%253A%2522vertical%2522%252C%2522id%2522%253A%2522DEVTOOLS%2522%252C%2522panels%2522%253A%255B%257B%2522type%2522%253A%2522PANEL%2522%252C%2522contentType%2522%253A%2522DEVTOOLS%2522%252C%2522id%2522%253A%2522clkolkcbj00b9356v1q1k3jh4%2522%257D%255D%252C%2522sizes%2522%253A%255B100%255D%257D%255D%252C%2522sizes%2522%253A%255B50%252C50%255D%257D%252C%2522tabbedPanels%2522%253A%257B%2522clkolkcbj00b7356v5kjcvj7e%2522%253A%257B%2522tabs%2522%253A%255B%257B%2522id%2522%253A%2522clkolkcbi00b6356v373f3ss3%2522%252C%2522mode%2522%253A%2522permanent%2522%252C%2522type%2522%253A%2522FILE%2522%252C%2522filepath%2522%253A%2522%252FREADME.md%2522%252C%2522state%2522%253A%2522IDLE%2522%257D%255D%252C%2522id%2522%253A%2522clkolkcbj00b7356v5kjcvj7e%2522%252C%2522activeTabId%2522%253A%2522clkolkcbi00b6356v373f3ss3%2522%257D%252C%2522clkolkcbj00b9356v1q1k3jh4%2522%253A%257B%2522id%2522%253A%2522clkolkcbj00b9356v1q1k3jh4%2522%252C%2522tabs%2522%253A%255B%255D%257D%252C%2522clkolkcbj00bb356v2r1kwyws%2522%253A%257B%2522tabs%2522%253A%255B%257B%2522id%2522%253A%2522clkolkcbj00ba356viaw09qcq%2522%252C%2522mode%2522%253A%2522permanent%2522%252C%2522type%2522%253A%2522TERMINAL%2522%252C%2522shellId%2522%253A%2522clkolkcr1000nfygi0bydahmn%2522%257D%255D%252C%2522id%2522%253A%2522clkolkcbj00bb356v2r1kwyws%2522%252C%2522activeTabId%2522%253A%2522clkolkcbj00ba356viaw09qcq%2522%257D%257D%252C%2522showDevtools%2522%253Atrue%252C%2522showShells%2522%253Atrue%252C%2522showSidebar%2522%253Atrue%252C%2522sidebarPanelSize%2522%253A15%257D)
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
