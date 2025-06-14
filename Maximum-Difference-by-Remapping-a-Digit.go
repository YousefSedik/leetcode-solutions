func minMaxDifference(num int) int {
	// Convert the number to a string to easily access its digits
	numStr := fmt.Sprintf("%d", num) // 11891
	i := 0
	for i = 0; i < len(numStr); i++ {
		if numStr[i] == '9' {
			continue
		} else {
			break
		}
	}
    if i == len(numStr){
		i--
	}
	maxStr := strings.ReplaceAll(numStr, string(numStr[i]), "9")
	minStr := strings.ReplaceAll(numStr, string(numStr[0]), "0")
	max, _ := strconv.ParseInt(maxStr, 10, 32)
	min, _ := strconv.ParseInt(minStr, 10, 32)
	return int(max - min)
}