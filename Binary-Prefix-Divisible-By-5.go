func prefixesDivBy5(nums []int) []bool {
    answer := make([]bool, len(nums))
    xI := 0 
    for i := 0; i < len(nums); i++{
        xI *= 2
        xI += nums[i]
        answer[i] = xI % 5 == 0
        xI %= 5
    }
    return answer
}