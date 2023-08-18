/**
* This program shows how a simple hash function works.
* The hash function size is 15. It can generate values of 0-14
* In the program the string is converted to byte array and the individual byte values
* are added. A modulo operation is performed on the cummulative sum and the result
* is retuned as hash value.
 */

package main

import (
	"fmt"
)

func main() {
	base := 15
	fmt.Println(getHash("John smith jhgjhguggugugugui", base))
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
