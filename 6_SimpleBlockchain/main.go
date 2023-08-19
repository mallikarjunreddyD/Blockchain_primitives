package main

import "fmt"

func main() {

	UTXOList = make(map[string]UTXO)
	//generateKeys()
	userData := loadKeys()
	genTransactions(userData)
}

func printList(UTXOList map[string]UTXO) {
	fmt.Println("-----------------------------------")
	for key, value := range UTXOList {
		fmt.Println(key, value)
	}
	fmt.Println("-----------------------------------")
}
