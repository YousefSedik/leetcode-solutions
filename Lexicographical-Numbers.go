func lexicalOrder(n int) []int {
	answer := make([]int, 0)
	var dfs func(start string)
	dfs = func(start string) {
		for i := range 10 {
			i, err := strconv.Atoi((start + strconv.Itoa(i)))
			if err == nil && i <= n {
				answer = append(answer, i)
				dfs(strconv.Itoa(i))
			} else {
				break
			}
		}
	}
	for i := 1; i < 10; i++ {
		if i <= n {
			answer = append(answer, i)
			dfs(strconv.Itoa(i))
		} else {
			break
		}
	}
	return answer
}