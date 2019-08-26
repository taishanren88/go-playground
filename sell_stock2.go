package main

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
	// Think in terms of profit
	// "sell", represents the profit if we sold on this day with a previous buy
	// "buy", represents the profit if we bought on this way using a previous sell
	// sell = max(previousSell, buy + price)
	// buy = max(previousBuy, previousSell - price)
	// Start from i =0;
	if len(prices) == 0 {
		return 0
	}
	buy := -prices[0]
	sell := 0
	for i := 1; i < len(prices); i++ {
		previousSell := sell
		sell = max(sell, buy+prices[i])
		buy = max(buy, previousSell-prices[i])
	}
	return sell
}
