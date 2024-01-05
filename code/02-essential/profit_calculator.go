package main

import "fmt"
import "os"

func main() {
	revenuCalculator()
	bank()
}

/* 
 Practices exercise
 Ask user for revenue, expenses, and tax rate
 Calculate
 Earn before tax, Earn after tax (profit) and ratio = Earn before tax / profit
 Print all result to console
 + Validate user input
 - non negative
 - not 0
 + Save result into file
*/

func revenuCalculator() {
	revenue := userInputValidation(askAndScanCalculator("Write your Revenue:"))
	expenses := userInputValidation(askAndScanCalculator("Write your Expenses:"))
	taxRate := userInputValidation(askAndScanCalculator("Write your Tax rate:"))
	
	earningBeforeTax, earningAfterTax, ratio := calculateEbtEatAndRatio(revenue, expenses, taxRate)

	saveDataToFile(earningBeforeTax, earningAfterTax, ratio)

	fmt.Printf("Earning before tax is: %.1f\n", earningBeforeTax)
	fmt.Printf("Earning after tax is: %.1f\n", earningAfterTax)
	fmt.Printf("Ratio is: %.1f\n", ratio)
}

func saveDataToFile(ebt, eat, ratio float64) {
	data := fmt.Sprintf("%0.2f\n%0.2f\n%0.2f", ebt, eat, ratio)
	os.WriteFile("profit.txt", []byte(data), 0644)
}

func userInputValidation(input float64) float64 {
	if input == 0 {
		panic("Input should be greater thant 0")
	} else if input < 0 {
		panic("Input should be positive")
	}
	return input
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