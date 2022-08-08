package stocks

func maxProfit(prices []int) int {
	minBuy := 1_000_000
	maxSell := -1

	for _, price := range prices {
		minBuy = min(minBuy, price)
		maxSell = max(maxSell, price-minBuy)
	}

	return maxSell
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
