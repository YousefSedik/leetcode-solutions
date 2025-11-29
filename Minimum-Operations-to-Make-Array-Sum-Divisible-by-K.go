1func minOperations(nums []int, k int) int {
2    sum := 0
3    for _, num := range(nums){
4        sum += num
5    }
6    return sum % k
7}