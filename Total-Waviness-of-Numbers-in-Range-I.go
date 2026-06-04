func totalWaviness(num1 int, num2 int) int {
    answer := 0
    for i := max(num1, 101); i <= num2; i++{
        i_str := strconv.Itoa(i)
        for j := 0; j < len(i_str) -2; j++{
            if i_str[j] < i_str[j+1] && i_str[j+1] > i_str[j+2] {
                answer++
            } else if i_str[j] > i_str[j+1] && i_str[j+1] < i_str[j+2] {
                answer++
            }
        }
    }
    return answer
}