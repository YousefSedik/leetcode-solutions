func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)

	for h := 0; h < 12; h++ {
		hBits := bits.OnesCount(uint(h))
		if hBits > turnedOn {
			continue
		}

		for m := 0; m < 60; m++ {
			if hBits+bits.OnesCount(uint(m)) == turnedOn {
				res = append(res, strconv.Itoa(h)+":"+fmt.Sprintf("%02d", m))
			}
		}
	}
	return res
}
