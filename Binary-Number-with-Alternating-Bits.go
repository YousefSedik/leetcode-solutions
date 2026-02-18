func hasAlternatingBits(n int) bool {
	binary_string := fmt.Sprintf("%b", n)
	for index, i := range []byte(binary_string) {
		if index % 2 == 0 && i != 49 {
			return false
		} else if index % 2 == 1 && i !=48{
			return false
		}
	}
	return true
}