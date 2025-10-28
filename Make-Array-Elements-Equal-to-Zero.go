func countValidSelections(nums []int) int {
	part1, part2, answer := 0, 0, 0
	for _, num := range nums {
		part2 += num
	}
	for _, num := range nums {
		if num == 0 {
			if part1 == part2 {
				answer += 2
			} else if part1 == part2-1 || part1 == part2+1 {
				answer += 1
			}
		} else {
			part1 += num
			part2 -= num
		}
	}
	return answer
}