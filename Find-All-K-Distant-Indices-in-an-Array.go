func findKDistantIndices(nums []int, key int, k int) []int {
    result := make([]int, 0)
    addedIndices := make(map[int]bool)
    for i := 0; i < len(nums) ; i++{
        if nums[i] == key{
            for j := i-1; j > max(i - k - 1, -1); j--{
                if addedIndices[j] == true{
                    continue
                }
                addedIndices[j] = true
                result = append(result, j)
            }
            if addedIndices[i] == false{
                addedIndices[i] = true
                result = append(result, i)
            }
            for j := i+1; j < min(i + k + 1, len(nums)); j++{
                if addedIndices[j] == true{
                    continue
                }
                addedIndices[j] = true
                result = append(result, j)
            }
        }
    }

    sort.Ints(result)
    return result
}