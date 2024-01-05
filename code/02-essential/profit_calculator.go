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
	revenue := askAndScan("Write your Revenue:")
	expenses := askAndScan("Write your Expenses:")
	taxRate := askAndScan("Write your Tax rate:")
	
	earningBeforeTax, earningAfterTax, ratio := calculateEbtEatAndRatio(revenue, expenses, taxRate)

	fmt.Print("Earning before tax is: ")
	fmt.Println(earningBeforeTax)

	fmt.Print("Earning after tax is: ")
	fmt.Println(earningAfterTax)

	fmt.Print("Ratio is: ")
	fmt.Println(ratio)
}

func askAndScan(ask string) float64 {
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