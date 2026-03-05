func minOperations(s string) int {
    option_1, option_2 := 0, 0
    flag := true 
    for _, bit := range(s){
        if flag{
            flag = false 
            if bit == '1'{
                option_1++
            } else {
                option_2++
            }
        } else {
            flag = true 
            if bit == '0'{
                option_1++
            } else {
                option_2++
            }
        }
    }    
    return min(option_1, option_2)     
}