func findLucky(arr []int) int {
    freq := [501]int{}
    answer := -1
    for i := 0; i < len(arr); i++{
        freq[arr[i]]++
    }
    for i := 1; i < 501; i++{
        if freq[i] == i{
            answer = i
        }
    }
    return answer
}