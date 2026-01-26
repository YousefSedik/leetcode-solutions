1func minimumAbsDifference(arr []int) [][]int {
2	sort.Ints(arr)
3	answer := make([][]int, 0)
4	minAbsDiff := math.MaxInt
5	for i := 0; i < len(arr)-1; i++ {
6		minAbsDiff = min(minAbsDiff, arr[i+1]-arr[i])
7	}
8	fmt.Println(minAbsDiff)
9	for i := 0; i < len(arr)-1; i++ {
10		if minAbsDiff == arr[i+1]-arr[i] {
11			answer = append(answer, []int{arr[i], arr[i+1]}) 
12		}
13	}
14	return answer
15}