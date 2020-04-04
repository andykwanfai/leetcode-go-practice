package main

func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	max := 0
	for i := 1; i < length; i++ {
		if profit := prices[i] - prices[i-1]; profit > 0 {
			max = profit + max
		}
	}
	return max
}
