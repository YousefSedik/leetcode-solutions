func maxAdjacentDistance(nums []int) int {
    answer := max(math.Abs(float64(nums[0] - nums[len(nums) - 1])))
    for i := 0; i < len(nums) - 1; i++{
        answer = max(answer, math.Abs(float64(nums[i+1] - nums[i])))
    }     
    return int(answer)
}