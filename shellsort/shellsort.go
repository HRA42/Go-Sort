package shellsort

func Sort(arr []int) []int {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			key := arr[i]
			j := i
			for j >= gap && arr[j-gap] > key {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = key
		}
	}
	return arr
}
