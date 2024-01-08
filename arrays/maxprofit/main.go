package maxprofit

// Find a window of increases where the difference is the largest
// Best day to sell will always be a peak
// Best day to buy will be the lowest point before
func maxProfit(prices []int) int {
	// Use a value which is out of bounds
	minPrice := prices[0]
	maxProfit := 0

	// Loop through prices once. O(n)
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if profit := prices[i] - minPrice; profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
