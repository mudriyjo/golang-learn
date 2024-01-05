package main

import "fmt"

func main() {
	age := 32
	
	agePointer := &age
	fmt.Println("Age: ", *agePointer)
	
	adultAge(*agePointer)
	fmt.Println("Age value: ", age)

	adultAgePointer(agePointer)
	fmt.Println("Age pointer: ", age)
}

func adultAgePointer(age *int) {
	*age = *age - 18
}

func adultAge(age int) {
	age = age - 18
}