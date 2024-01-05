package main

import "fmt"
import "example.com/bank/utils"
import "github.com/Pallinder/go-randomdata"

const balanceFileName = "balance.txt"

func main() {
	balance, err := utils.ReadIntFromFile(balanceFileName, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Hello to Go Bank!")
	fmt.Println("We open 24/7, you can connect: ", randomdata.PhoneNumber())
	
	for {
		printMenu()
		choice := utils.AskAndScan("What's your choice: ")
		
		switch choice {
			case 1:
				printBalance(balance)
			case 2:
				moneyForDeposit := utils.AskAndScan("How much money do you want to deposit: ")
				
				if moneyForDeposit < 0 {
					fmt.Println("Amount should be greater than 0.")
					continue
				}

				balance += moneyForDeposit
				utils.WriteIntToFile(balanceFileName, balance)
				printBalance(balance)
			case 3:
				moneyForWithdraw := utils.AskAndScan("How much money do you want to withdraw: ")
				
				if moneyForWithdraw < 0 {
					fmt.Println("Amount should be greater than 0.")
					continue
				}

				if (balance < moneyForWithdraw) {
					fmt.Println("You don't have enough money.")
					printBalance(balance)
				} else {
					balance -= moneyForWithdraw
					utils.WriteIntToFile(balanceFileName, balance)
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