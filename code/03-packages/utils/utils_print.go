package utils

import "fmt"

func AskAndScan(ask string) int {
	var result int
	fmt.Print(ask)
	fmt.Scan(&result)
	return result
}