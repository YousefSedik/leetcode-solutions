func maxAdjacentDistance(nums []int) int {
    // absolute(value) = max(value, value * -1)
    answer := max(nums[0] - nums[len(nums) - 1], (nums[0] - nums[len(nums) - 1]) * -1)
    for i := 0; i < len(nums) - 1; i++{
        answer = max(answer, nums[i+1] - nums[i], (nums[i+1] - nums[i]) * -1)
    }     
    return answer
}