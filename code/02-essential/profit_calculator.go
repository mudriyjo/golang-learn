package main

import "fmt"

/* 
 Practices exercise
 Ask user for revenue, expenses, and tax rate
 Calculate
 Earn before tax, Earn after tax (profit) and ratio = Earn before tax / profit
 Print all result to console
*/
func main() {
	revenuCalculator()
	bank()
}

func revenuCalculator() {
	revenue := askAndScanCalculator("Write your Revenue:")
	expenses := askAndScanCalculator("Write your Expenses:")
	taxRate := askAndScanCalculator("Write your Tax rate:")
	
	earningBeforeTax, earningAfterTax, ratio := calculateEbtEatAndRatio(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax is: %.1f\n", earningBeforeTax)
	fmt.Printf("Earning after tax is: %.1f\n", earningAfterTax)
	fmt.Printf("Ratio is: %.1f\n", ratio)
}

func askAndScanCalculator(ask string) float64 {
	var result float64
	fmt.Print(ask)
	fmt.Scan(&result)
	return result
}

func calculateEbtEatAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := (revenue  - expenses) * (1 - taxRate / 100)
	ratio := earningBeforeTax / earningAfterTax
	return earningBeforeTax, earningAfterTax, ratio
}