package main

func main() {

}
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		var j = i + 1
		for j < len(prices) && prices[j] > prices[i] {
			j++
		}
		if j < len(prices) {
			prices[i] -= prices[j]
		}
	}
	return prices
}
