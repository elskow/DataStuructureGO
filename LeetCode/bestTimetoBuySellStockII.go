package main

func maxProfit(prices []int) int {
	var (
		maxProfit int
	)
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))

	prices = []int{1, 2, 3, 4, 5}
	println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	println(maxProfit(prices))
}
