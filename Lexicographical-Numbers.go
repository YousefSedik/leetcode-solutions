func lexicalOrder(n int) []int {
	answer := make([]int, n)
	index := 0
	var dfs func(start string)
	dfs = func(start string) {
		for i := range 10 {
			i, _ := strconv.Atoi((start + strconv.Itoa(i)))
			if i <= n {
				answer[index] = i
				index++
				dfs(strconv.Itoa(i))
			} else {
				break
			}
		}
	}
	for i := 1; i < 10; i++ {
		if i <= n {
			answer[index] = i
			index++
			dfs(strconv.Itoa(i))

		} else {
			break
		}
	}
	return answer
}