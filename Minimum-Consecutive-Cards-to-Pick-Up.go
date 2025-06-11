type Pair struct {
	current int
	next    int
}

func minimumCardPickup(cards []int) int {
	mapper := make(map[int]Pair)
	answer := 999999
	for index, value := range cards {
		index += 1
		if mapper[value].current == 0 {
			// set the current, compute the min
			mapper[value] = Pair{current: index, next: 0}
		} else if mapper[value].next == 0 {
			// set the next, compute the min
			mapper[value] = Pair{current: mapper[value].current, next: index}
			answer = min(answer, mapper[value].next-mapper[value].current+1)
		} else {
			// update current, compute the min
			mapper[value] = Pair{current: mapper[value].next, next: index}
			answer = min(answer, mapper[value].next-mapper[value].current+1)
		}
	}
	if answer == 999999 {
		return -1
	}
	return answer
}