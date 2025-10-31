func getSneakyNumbers(nums []int) []int {
    alreadyExist := make(map[int]bool)
    result := make([]int, 0, 2) 
    for i := 0; i < len(nums); i++{
        if alreadyExist[nums[i]] == true{
            result = append(result, nums[i])
            if len(result) == 2{
                return result
            }
        } else {
            alreadyExist[nums[i]] = true
        }
    }   
    return result
}