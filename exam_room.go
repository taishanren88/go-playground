package main

type ExamRoom struct {
}

func Constructor(N int) ExamRoom {
	// Create a set of available intervals and have it be sorted by distance, where the top is the one with the largest distance
	// Whenever someone comes in , remove the first entry from the set and gets it seat
	// initially add [-1, N] as the first interval
	// Whenever a seat is taken split it into two intervals :
	// If seat is 0 , split into [-1, 0] and [0, end]
	// If seat is end, split into [start, N-1], [N-1, end]
	// Else split into [0, mid], [mid, end]
	// During leave(), find the interval which contains this seat as tail and head
	// [start, target], [target, end]. Remove those two and merge them back.
}

func (this *ExamRoom) Seat() int {
	return -1
}

func (this *ExamRoom) Leave(p int) {
	return -1
}
