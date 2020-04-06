package main

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 0; i < length; i++ {
		profit := prices[i] - min
		if profit > max {
			max = profit
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}
