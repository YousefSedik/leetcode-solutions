func mincostTickets(days []int, costs []int) int {
	daySet := make(map[int]bool)
	for _, d := range days {
		daySet[d] = true
	}

	memo := make(map[int]int)
	var dp func(int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if val, ok := memo[day]; ok {
			return val
		}
		if !daySet[day] {
			memo[day] = dp(day + 1)
			return memo[day]
		}
		memo[day] = min(
			costs[0]+dp(day+1),
			costs[1]+dp(day+7),
			costs[2]+dp(day+30),
		)
		return memo[day]
	}
	return dp(days[0])
}
