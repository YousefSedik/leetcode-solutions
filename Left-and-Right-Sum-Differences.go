func leftRightDifference(nums []int) []int {
    n := len(nums)
    answer := make([]int, n)
    leftSum, rightSum := 0, 0

    for i := range(n){
        leftSum += nums[i]
        rightSum += nums[n-i-1]
        answer[i] += leftSum
        answer[n-i-1] -= rightSum
    }

    for i := range(n){
        answer[i] = int(math.Abs(float64(answer[i])))
    }
    
    return answer
}