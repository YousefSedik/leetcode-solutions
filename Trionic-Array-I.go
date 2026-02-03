1func isTrionic(nums []int) bool {
2	decIndexStart, decIndexEnd := -1, -1 
3	for i := 0; i < len(nums)-1; i++ {
4		if nums[i] >= nums[i+1]{
5			decIndexStart = i
6			break
7		}
8	}
9	if decIndexStart == 0 || decIndexStart == len(nums) - 1 || decIndexStart == -1 {
10		return false
11	}
12	for i := len(nums) -1 ; i > 0; i-- {
13		if nums[i-1] >= nums[i]{
14			decIndexEnd = i
15			break
16		}
17	}
18	if decIndexEnd == len(nums) - 1 {
19		return false
20	}
21	for i := decIndexStart; i < decIndexEnd; i++ {
22		if nums[i] <= nums[i+1]{
23			return false
24		}
25	}
26	return true
27}