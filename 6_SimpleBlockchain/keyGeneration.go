package main

import (
	"encoding/json"
	"fmt"
	"os"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

type user struct {
	name       string
	address    string
	publicKey  string
	privateKey string
}

var users map[int]user

func generateKeys() {
	users = make(map[int]user)
	userNames := [7]string{"Alice", "Bob", "Charlie", "Darth", "Earl", "Fin", "Grace"}
	for index, name := range userNames {
		address, publicKey, privateKey := genKey()
		user := user{
			name:       name,
			address:    address,
			publicKey:  publicKey,
			privateKey: privateKey,
		}
		users[index] = user
	}
	printKeys()
	saveKeys()
}

func genKey() (string, string, string) {
	key, err := ethCrypto.GenerateKey()

	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	publicKeyBytes := ethCrypto.FromECDSAPub(&key.PublicKey)
	publicKey := ethCommon.Bytes2Hex(publicKeyBytes)
	address := ethCrypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := ethCommon.Bytes2Hex(ethCrypto.FromECDSA(key))
	return address, publicKey, privateKey
}
func printKeys() {
	fmt.Println("The generated keys are")
	for index, value := range users {
		fmt.Println(index, value)
		fmt.Println("----------------------------")
	}
}

func saveKeys() {
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Marshal the map to JSON and write to the file
	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
