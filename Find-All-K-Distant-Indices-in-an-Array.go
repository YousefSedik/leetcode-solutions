func findKDistantIndices(nums []int, key int, k int) []int {
	keyIndices := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			keyIndices = append(keyIndices, i)
		}
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if keyIndices[j]+k < i && j != len(keyIndices) - 1{
			j++
		}
		if i >= keyIndices[j]-k && i <= keyIndices[j]+k{
			result = append(result, i)
		}
		if j != len(keyIndices) - 1 && keyIndices[j+1] == i {
			j++
		}
	}

	return result
}