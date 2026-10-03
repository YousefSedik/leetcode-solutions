func longestValidParentheses(s string) int {
    i, l_count, r_count, longest := 0, 0, 0, 0

    for i < len(s){
        if s[i] == '('{
            l_count++
        } else {
            r_count++
        }

        if l_count == r_count{
            longest = max(longest, l_count + r_count)
        } else if r_count > l_count {
            r_count, l_count = 0, 0
        }
        i++
    }   

    i = len(s) - 1
    l_count, r_count = 0, 0
    
    for i >= 0 {
        if s[i] == '('{
            l_count++
        } else {
            r_count++
        }

        if l_count == r_count{
            longest = max(longest, l_count + r_count)
        } else if r_count < l_count {
            r_count, l_count = 0, 0
        }
        i--
    }   

    return longest
}