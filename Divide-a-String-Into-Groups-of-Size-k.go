func divideString(s string, k int, fill byte) []string {
	answer := make([]string, 0, len(s)/k+1)
	stop, index := false, 0
	sw := strings.Builder{}

	for stop == false {
		if index < len(s) {
			sw.WriteByte(s[index])
			index++
		} else {
			sw.WriteByte(fill)
		}
		if len(sw.String()) == k {
			answer = append(answer, sw.String())
			sw.Reset()
			if index >= len(s) {
				stop = true
			}
		}
	}
	return answer
}