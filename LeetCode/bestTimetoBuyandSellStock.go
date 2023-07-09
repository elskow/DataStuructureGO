package main

func maxProfit(prices []int) int {
	var (
		minPrice  int = 1<<31 - 1
		maxProfit int
	)
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	println(maxProfit(prices))
}
