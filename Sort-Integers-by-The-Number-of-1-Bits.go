func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		b1 := bits.OnesCount(uint(arr[i]))
		b2 := bits.OnesCount(uint(arr[j]))
		if b1 == b2 {
			return arr[i] < arr[j]
		}
		return b1 < b2
	})
	return arr
}