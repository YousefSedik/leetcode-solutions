func longestBalanced(s string) int {
    localLongest, globalLongest := 1, 1
    
    for i := 0; i < len(s); i++{
        charCounter := make(map[byte]int)
        count, countCount := 0, 0
        for j := i; j < len(s); j++{
            charCounter[s[j]]++
            if charCounter[s[j]] > count{
                countCount = 1
                count = charCounter[s[j]]
            } else if charCounter[s[j]] == count {
                countCount++
            }
            if countCount == len(charCounter){
                localLongest = j - i + 1
            }
        }
        globalLongest = max(globalLongest, localLongest)
    }
    return globalLongest
}