func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	leftProfit := make([]int, n)
	rightProfit := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
		leftProfit[i] = int(math.Max(float64(leftProfit[i-1]), float64(prices[i]-minPrice)))
	}
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxPrice = int(math.Max(float64(maxPrice), float64(prices[i])))
		rightProfit[i] = int(math.Max(float64(rightProfit[i+1]), float64(maxPrice-prices[i])))
	}
	maxProfit := 0
	for i := 0; i < n; i++ {
		maxProfit = int(math.Max(float64(maxProfit), float64(leftProfit[i]+rightProfit[i])))
	}

	return maxProfit
}