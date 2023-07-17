package bubblesort

func Sort(arr []int) []int {
	for i := range arr {
		for j := range arr[:len(arr)-i-1] {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
