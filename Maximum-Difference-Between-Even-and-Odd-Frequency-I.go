func maxDifference(s string) int {
    freq := [27]int8{}
    for _, letter := range(s){
        freq[int(letter) - 97] += 1
    }
    var a1, a2 int8 = 0, 100
    for _, occ := range(freq){
        if occ != 0 {
            if occ % 2 == 0{
                a2 = min(a2, occ)
            } else {
                a1 = max(a1, occ)
            }
        }
    }
    return int(a1 - a2)
    
}
