package main

import "fmt"

func main() {
	fmt.Println("price calculator")
	var prices []float64 = []float64{10, 20, 30}
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.2}


	results := make(map[float64][]float64)

	fmt.Println(prices)
	fmt.Println(results)

	for _, taxRate := range taxRates {
		var taxIncludingPrices []float64 = make([]float64, leb(prices))
		fmt.Println(taxRate)
	} 

	for _, result := range results {
		fmt.Println(result)
	}
}