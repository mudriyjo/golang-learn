package utils

import "os"
import "errors"
import "strconv"
import "fmt"

const fileRight = 0644

func ReadIntFromFile(fileName string, defaultInt int) (int, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultInt, errors.New("file not found")
	}

	textInt := string(data)
	intResult, err := strconv.ParseInt(textInt, 10, 64)
	
	if err != nil {
		return defaultInt, errors.New("can't parse file to int")
	}

	return int(intResult), nil
}

func WriteIntToFile(fileName string, value int) {
	intText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(intText), fileRight)
}