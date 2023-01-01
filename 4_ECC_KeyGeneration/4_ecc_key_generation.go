package main

import (
	"fmt"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, err := ethCrypto.GenerateKey()

	if err != nil {
		fmt.Println(err)
		return
	}
	publicKeyBytes := ethCrypto.FromECDSAPub(&key.PublicKey)
	publicKey := ethCommon.Bytes2Hex(publicKeyBytes)
	address := ethCrypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := ethCommon.Bytes2Hex(ethCrypto.FromECDSA(key))
	fmt.Println("PublicKey: ", publicKey)
	fmt.Println("Address: ", address)
	fmt.Println("PrivateKey: ", privateKey)

}
