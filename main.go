package main

import "fmt"

func main() {
	fmt.Println("price calculator")
	var prices []float64 = []float64{10, 20, 30}
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.2}


	result := make(map[float64][]float64)

	fmt.Println(prices)
	fmt.Println(result)

	for _, taxRate := range taxRates {
		fmt.Println(taxRate)
	} 
}