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
	var revenue, expenses, taxRate float64

	fmt.Print("Write your Revenue:")
	fmt.Scan(&revenue)

	fmt.Print("Write your Expenses:")
	fmt.Scan(&expenses)

	fmt.Print("Write your Tax rate:")
	fmt.Scan(&taxRate)
	
	earningBeforeTax := revenue - expenses
	earningAfterTax := (revenue  - expenses) * (1 - taxRate / 100)
	ratio := earningBeforeTax / earningAfterTax

	fmt.Print("Earning before tax is: ")
	fmt.Println(earningBeforeTax)

	fmt.Print("Earning after tax is: ")
	fmt.Println(earningAfterTax)

	fmt.Print("Ratio is: ")
	fmt.Println(ratio)
}