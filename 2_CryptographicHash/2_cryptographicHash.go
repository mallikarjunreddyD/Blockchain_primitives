/** A program to demonstrate differnt commonly used hash functions*/

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/wealdtech/go-merkletree/keccak256"
)

func main() {
	fmt.Println(getHash("John smiti"))
	//fmt.Println(getHash("Lisa Smith"))
	//fmt.Println(getHash("Sam Doe"))
	//fmt.Println(getHash("Sandra Dee"))
}
func getHash(data string) string {
	result := fmt.Sprintf("MD5: %x, SHA256: %x, Keccak256: %x",
		md5.Sum([]byte(data)),
		sha256.Sum256([]byte(data)), keccak256.New().Hash([]byte(data)))
	return result
}
