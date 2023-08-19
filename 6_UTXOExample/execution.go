package main

import (
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func execute(value string, owner string) string {
	trans := UTXO{}
	transId := getTransactionId(value, owner)
	trans.value = value
	trans.owner = owner
	trans.spent = false
	UTXOList[transId] = trans
	return transId
}

/*
To calculate transactionId, the full input
and output scripts must be considered.
For simplicity, this function takes only value and owner
*/
func getTransactionId(value string, owner string) string {
	Id := ethCrypto.Keccak256Hash([]byte(owner), []byte(value)).Hex()
	Id = Id[:10]
	return Id
}
func executeTransaction(inputs []input, outputs []UTXO) bool {
	for _, val := range inputs {
		result := verifySignature(val)
		if result == false {
			return false
		}
		trans := UTXOList[val.Id]
		trans.spent = true
		UTXOList[val.Id] = trans
	}
	for _, val := range outputs {
		execute(val.value, val.owner)
	}
	return true
}
