func maxDifference(s string) int {
    freq := [27]int{}
    for _, letter := range(s){
        freq[int(letter) - 97] += 1
    }
    a1, a2 := -1, 9999
    for _, occ := range(freq){
        if occ != 0 {
            if occ % 2 == 0{
                a2 = min(a2, occ)
            } else {
                a1 = max(a1, occ)
            }
        }
    }
    return a1 - a2
    
}
