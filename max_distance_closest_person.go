package main

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
func maxDistToClosest(seats []int) int {
	// Consider distance from the left separately. it is just (index)
	// Consider distance from the right separately, it is jsut (size - 1 - index)
	// Consider everything in between as (Right - left) / 2
	N := len(seats)
	maxDist := 0
	start := 0
	for ; start < N && seats[start] == 0; start++ {

	}
	if start < N {
		maxDist = start
	}
	for i := start + 1; i < N; i++ {
		if seats[i] == 1 {
			maxDist = max((i-start)/2, maxDist)
			start = i
		}
	}

	maxDist = max(N-start-1, maxDist)
	return maxDist
}
