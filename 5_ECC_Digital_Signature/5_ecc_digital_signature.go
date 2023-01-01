package main

import (
	"fmt"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	message := []byte("Hello")
	publicKey, signature := genSignature(message)
	fmt.Println(signature)
	publicKey1 := verifySignature(message, signature)
	if ethCommon.Bytes2Hex(publicKey) == ethCommon.Bytes2Hex(publicKey1) {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

}
func genSignature(messgae []byte) ([]byte, string) {
	key, err := ethCrypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}
	publicKey := ethCrypto.FromECDSAPub(&key.PublicKey)
	privateKey := ethCommon.Bytes2Hex(ethCrypto.FromECDSA(key))
	priv, err := ethCrypto.HexToECDSA(privateKey)

	hash := ethCrypto.Keccak256Hash(messgae)
	if err != nil {
		fmt.Println(err)
	}
	signatureBytes, err := ethCrypto.Sign(hash.Bytes(), priv)
	if err != nil {
		fmt.Println(err)
	}
	signature := ethCommon.Bytes2Hex(signatureBytes)
	return publicKey, signature
}
func verifySignature(message []byte, signature string) []byte {
	publickey, err := ethCrypto.Ecrecover(ethCrypto.Keccak256Hash(message).Bytes(), ethCommon.Hex2Bytes(signature))
	if err != nil {
		fmt.Println(err)
	}
	return publickey
}
