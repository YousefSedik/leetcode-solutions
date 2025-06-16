func maximumDifference(nums []int) int {
    answer := -1
    MinSoFar := int(^uint(0)  >> 1)  
    for i := 0; i < len(nums); i++{
        MinSoFar = min(MinSoFar, nums[i])
        if i > 0 && nums[i] > MinSoFar {
            answer = max(answer,nums[i] - MinSoFar)
        }
    }
    return answer
}