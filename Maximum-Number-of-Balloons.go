func maxNumberOfBalloons(text string) int {
    countA, countB, countL, countO, countN := 0, 0, 0, 0, 0
    for i := range(len(text)){
        if text[i] == 'b'{
            countB++
        } else if text[i] == 'a'{
            countA++
        } else if text[i] == 'l'{
            countL++
        } else if text[i] == 'o'{
            countO++
        } else if text[i] == 'n'{
            countN++
        }
    }
    
    return min(countA, countB, countL/2, countO/2, countN)
}