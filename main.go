package main

import (
	"fmt"
)

func main() {
	base := 15
	fmt.Println(getHash("John smith", base))
	fmt.Println(getHash("Lisa Smith", base))
	fmt.Println(getHash("Sam Doe", base))
	fmt.Println(getHash("Sandra Dee", base))
}
func getHash(data string, b int) int {
	decimalValue := fmt.Sprintf("%d", []byte(data))

	sum := 0
	for _, val := range decimalValue {
		sum = sum + int(byte(val))

	}
	return sum % b
}
