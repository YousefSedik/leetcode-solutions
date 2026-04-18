func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func mirrorDistance(n int) int {
	reversedN, err := strconv.ParseInt(Reverse(strconv.FormatInt(int64(n), 10)), 10, 64)
	if err != nil {
        return n
	}
    return AbsInt(n - int(reversedN))
}