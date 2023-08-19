package main

import "fmt"

func main() {

	UTXOList = make(map[string]UTXO)
	//generateKeys()
	userData := loadKeys()
	GenesisTransactions(userData)
	Transaction1(userData)
	Transaction2(userData)
	Transaction3(userData)
}

func printList(UTXOList map[string]UTXO) {
	fmt.Println("-----------------------------------")
	for key, value := range UTXOList {
		fmt.Println(key, value)
	}
	fmt.Println("-----------------------------------")
}
