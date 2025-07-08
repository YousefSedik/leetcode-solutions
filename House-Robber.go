func rob(nums []int) int {
	// 2 options, rob current and skip, skip
	mem := make(map[int]int)
    var solve func(i int) int
	solve = func(i int) int {
		if i >= len(nums) {
			return 0
		}
        if v, ok := mem[i]; ok{
            return v
        }
        mem[i] = max(solve(i+1), nums[i]+ solve(i+2))
		return mem[i]
	}
    return max(solve(0), solve(1))
}