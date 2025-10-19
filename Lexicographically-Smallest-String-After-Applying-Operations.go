func findLexSmallestString(s string, a, b int) string {
	visited := map[string]bool{s: true}
	queue := []string{s}
	answer := s
	for head := 0; head < len(queue); head++ {
		v := queue[head]
		if v < answer {
			answer = v
		}
		bytesArr := []byte(v)
		for i := 1; i < len(bytesArr); i += 2 {
			bytesArr[i] = byte((int(bytesArr[i]-'0') + a) % 10 + '0')
		}
		addStr := string(bytesArr)
		if !visited[addStr] {
			visited[addStr] = true
			queue = append(queue, addStr)
		}

		rotated := v[len(v)-b:] + v[:len(v)-b]
		if !visited[rotated] {
			visited[rotated] = true
			queue = append(queue, rotated)
		}
	}
	return answer
}
