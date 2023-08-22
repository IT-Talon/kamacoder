package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	var profit float64
	min := math.MaxFloat64
	for _, price := range prices {
		min = math.Min(min, float64(price))
		profit = math.Max(float64(price)-min, profit)
	}
	return int(profit)
}

func findMaxPrice(prices []int) int {
	max := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		}
	}
	return max
}
