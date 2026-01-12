1func abs(x int) int {
2    if x < 0 {
3        return -x
4    }
5    return x
6}
7
8func minTimeToVisitAllPoints(points [][]int) int {
9	minTime := 0
10	for i := 0; i < len(points)-1; i++ {
11		dx := abs(points[i][0] - points[i+1][0])
12		dy := abs(points[i][1] - points[i+1][1])
13
14		if dx > dy {
15			minTime += dx
16		} else {
17			minTime += dy
18		}
19	}
20	return minTime
21}