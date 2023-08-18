package main

import (
	"fmt"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

/** UTXO structure */
type UTXO struct {
	value string
	owner string
	spent bool
}

/** Transaction structure */
type transaction struct {
	inputs  []UTXO
	outputs []UTXO
}

/** A map to store all the UTXO's **/
var UTXOList map[string]UTXO

func genTransactions() {
	UTXOList = make(map[string]UTXO)
	addTrans("100", users[0].address) // Genesis transaction
	addTrans("50", users[0].address)  // Genesis transaction
	addTrans("70", users[1].address)  // Genesis transaction
	addTrans("60", users[2].address)  // Genesis transaction

	fmt.Println("UTXOList after Genesis") // UTXOList after genesis transactions
	printList(UTXOList)

	/* Trans1Inputs := []string{"0xe12230b9", "0x069884ac"} // Transaction 1 inputs
	Trans1Outputs := []UTXO{{                            // Transaction 1 outputs
		value: "120",
		owner: "Darth",
	}, {
		value: "30",
		owner: "Alice",
	}}

	executeTransaction(Trans1Inputs, Trans1Outputs) // Execute transaction 1
	fmt.Println("UTXOList after Transaction 1")
	printList(UTXOList) // UTXOList after transaction 1

	Trans2Inputs := []string{"0xb8a3b7f7"} // Transaction 1 inputs
	Trans2Outputs := []UTXO{{              // Transaction 1 outputs
		value: "50",
		owner: "Earl",
	}, {
		value: "30",
		owner: "Bob",
	},
		{
			value: "30",
			owner: "Fin",
		}}

	executeTransaction(Trans2Inputs, Trans2Outputs) // Execute transaction 1
	fmt.Println("UTXOList after Transaction 2")
	printList(UTXOList) // UTXOList after transaction 2

	Trans3Inputs := []string{"0xc664105e", "0x79fe7899"} // Transaction 1 inputs
	Trans3Outputs := []UTXO{{                            // Transaction 1 outputs
		value: "120",
		owner: "Darth",
	}, {
		value: "",
		owner: "Alice",
	}}

	executeTransaction(Trans3Inputs, Trans3Outputs) // Execute transaction 1
	fmt.Println("UTXOList after Transaction 3")
	printList(UTXOList) // UTXOList after transaction  */
}

func addTrans(value string, owner string) string {
	trans := UTXO{}
	trans.value = value
	trans.owner = owner
	Id := ethCrypto.Keccak256Hash([]byte(owner), []byte(value)).Hex()
	Id = Id[:10]
	trans.spent = false
	UTXOList[Id] = trans
	return Id
}
func executeTransaction(inputs []string, outputs []UTXO) {
	validateInputs()
	validateOutputs()
	generateTransaction(inputs, outputs)
}
func validateInputs() {
	verifyOwnership()
}
func generateTransaction(inputs []string, outputs []UTXO) {
	for _, val := range inputs {
		trans := UTXOList[val]
		trans.spent = true
		UTXOList[val] = trans
	}
	for _, val := range outputs {
		addTrans(val.value, val.owner)
	}
}

func printList(UTXOList map[string]UTXO) {
	fmt.Println("-----------------------------------")
	for key, value := range UTXOList {
		fmt.Println(key, value)
	}
	fmt.Println("-----------------------------------")
}
func validateOutputs() {}
func verifyOwnership() {}
func genSignature(privateKey string, data string) string {
	hash := ethCrypto.Keccak256Hash([]byte(data))

	signatureBytes, err := ethCrypto.Sign(hash.Bytes(), ethCrypto.ToECDSAUnsafe((ethCommon.Hex2Bytes(privateKey))))
	if err != nil {
		fmt.Println(err)
	}
	signature := ethCommon.Bytes2Hex(signatureBytes)
	return signature
}
