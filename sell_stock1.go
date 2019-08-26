package main

import "math"

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxProfit(prices []int) int {
	// Find the minimum so far value
	// At each point , just compute current price - that point
	minPrice := math.MaxInt32
	maxProfit := 0
	for _, price := range prices {
		minPrice = min(price, minPrice)
		profit := price - minPrice
		maxProfit = max(profit, maxProfit)
	}
	return maxProfit
}
