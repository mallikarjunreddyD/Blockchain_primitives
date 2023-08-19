package main

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func genSignature(privateKey string, data string) string {
	hash := ethCrypto.Keccak256Hash([]byte(data))
	signatureBytes, err := ethCrypto.Sign(hash.Bytes(), ethCrypto.ToECDSAUnsafe((ethCommon.Hex2Bytes(privateKey))))
	if err != nil {
		return ""
	}
	signature := ethCommon.Bytes2Hex(signatureBytes)
	return signature
}
func verifySignature(in input) bool {
	publickey, err := ethCrypto.Ecrecover(ethCrypto.Keccak256Hash([]byte(in.Id)).Bytes(), ethCommon.Hex2Bytes(in.signature))
	if err != nil {
		return false
	}
	publicKeyECDSA, err := ethCrypto.UnmarshalPubkey(publickey)
	if err != nil {
		return false
	}
	address := ethCrypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if address == UTXOList[in.Id].owner {
		return true
	}
	return false
}
