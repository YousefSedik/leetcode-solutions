func minCost(colors string, neededTime []int) int {
	res := 0
	for i := 0; i < len(colors)-1; i++ {
		partial_sum := 0
		max_value := 0
		if colors[i] == colors[i+1] {
			for i != len(colors)-1 && colors[i] == colors[i+1]{
				partial_sum += neededTime[i]
				max_value = max(max_value, neededTime[i])
				i++
			}
			if i != len(colors){
				partial_sum += neededTime[i]
				max_value = max(max_value, neededTime[i])
				// i++
			}
			partial_sum -= max_value
			res += partial_sum
		}
	}
	return res
}