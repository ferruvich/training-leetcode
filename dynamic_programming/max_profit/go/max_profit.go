package max_profit

import "sort"

func MaxProfit(prices []int) int {
	return MaxProfitLinear(prices)
}

// O(n^2) time, as we're traversing the list for each element
// O(1) space, as we're only using one variable
func MaxProfitQuadratic(prices []int) int {
	var max int
	for i, buy := range prices {
		for _, sell := range prices[i:] {
			if sell > buy && max < sell-buy {
				max = sell - buy
			}
		}
	}

	return max
}

// O(n^2logn) time as we're sorting the remaining list, for each element
// O(n^2) space, as we're using a list for each element
func MaxProfitSuccessiveSorting(prices []int) int {
	var max int
	for i, buy := range prices {
		k := append([]int{}, prices[i:]...)
		sort.Ints(k)
		if k[len(k)-1]-buy > max {
			max = k[len(k)-1] - buy
		}
	}

	return max
}

// O(n) time, as we're traversing the list once
// O(1) space, as we're only using two variables
func MaxProfitLinear(prices []int) int {
	var minPrice, maxProfit int
	minPrice = prices[0]
	for _, el := range prices[1:] {
		if el < minPrice {
			minPrice = el
			continue
		}
		if el-minPrice > maxProfit {
			maxProfit = el - minPrice
		}
	}

	return maxProfit
}
