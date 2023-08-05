package insertionsort

func Sort(arr []int) []int {
	for i, v := range arr {
		temp := v
		j := i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}
