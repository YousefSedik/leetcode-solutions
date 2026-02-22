func binaryGap(n int) int {
	bs := strconv.FormatInt(int64(n), 2)
	x, y, gab := -1, -1, 0
	for index, i := range(bs) {
		if i == '1' {
			if x == -1 {
				x = index
			} else {
				y = x
				x = index
				gab = max(gab, x-y)
			}
		}
	}
	return gab
}