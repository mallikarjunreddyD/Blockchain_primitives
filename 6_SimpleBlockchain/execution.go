package main

import (
	"fmt"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func execute(value string, owner string) string {
	trans := UTXO{}
	trans.value = value
	trans.owner = owner
	Id := ethCrypto.Keccak256Hash([]byte(owner), []byte(value)).Hex()
	Id = Id[:10]
	trans.spent = false
	UTXOList[Id] = trans
	return Id
}

func executeTransaction(inputs []input, outputs []UTXO) {
	for _, val := range inputs {
		trans := UTXOList[val.Id]
		trans.spent = true
		UTXOList[val.Id] = trans
	}
	for _, val := range outputs {
		execute(val.value, val.owner)
	}
}

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
