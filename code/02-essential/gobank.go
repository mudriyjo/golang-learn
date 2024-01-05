package main

import "fmt"
import "os"
import "errors"
import "strconv"

const balanceFileName = "balance.txt"
const fileRight = 0644

func readBalanceFromFile() (int, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 1000, errors.New("balance file not found")
	}

	textBalance := string(data)
	balance, err := strconv.ParseInt(textBalance, 10, 64)
	
	if err != nil {
		return 1000, errors.New("can't parse balance file to int")
	}

	return int(balance), nil
}

func writeBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), fileRight)
}

func bank() {
	balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Hello to Go Bank!")

	for {
		printMenu()
		choice := askAndScan("What's your choice: ")
		
		switch choice {
			case 1:
				printBalance(balance)
			case 2:
				moneyForDeposit := askAndScan("How much money do you want to deposit: ")
				
				if moneyForDeposit < 0 {
					fmt.Println("Amount should be greater than 0.")
					continue
				}

				balance += moneyForDeposit
				writeBalanceToFile(balance)
				printBalance(balance)
			case 3:
				moneyForWithdraw := askAndScan("How much money do you want to withdraw: ")
				
				if moneyForWithdraw < 0 {
					fmt.Println("Amount should be greater than 0.")
					continue
				}

				if (balance < moneyForWithdraw) {
					fmt.Println("You don't have enough money.")
					printBalance(balance)
				} else {
					balance -= moneyForWithdraw
					writeBalanceToFile(balance)
					printBalance(balance)
				}
			default:
				fmt.Println("Goodbye!")
				return
		}
	}
}

func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func printBalance(balance int) {
	fmt.Printf("You balance is: %v\n", balance)
}

func askAndScan(ask string) int {
	var result int
	fmt.Print(ask)
	fmt.Scan(&result)
	return result
}