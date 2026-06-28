func intAbs(n int) int{
	if n < 0 {
		return -1 * n
	}
	return n
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	slices.Sort(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if intAbs(arr[i] - arr[i-1]) > 1{
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}