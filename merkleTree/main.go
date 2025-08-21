package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
)

const n int = 8
const twopn int = 2*n - 1

type MerkleTree struct {
	digest [twopn]string
}

func main() {

	leafs := [n]string{"0", "1", "2", "3", "4", "5", "6", "7"}
	M := setLeaf(leafs)
	M = build_tree(M)
	print(M)
	requestedLeaves := []int{0}
	proof := compressedProof(M, requestedLeaves)
	fmt.Println(len(proof))
	root := M.digest[0]
	result := verifyMerkle(root, 1, []string{M.digest[8]}, []int{7}, proof)
	fmt.Println(result)
}

func setLeaf(leafs [8]string) MerkleTree {
	var M MerkleTree
	hash := sha256.New()
	for i, leaf := range leafs {
		hash.Write([]byte(leaf))
		M.digest[i+n-1] = string(hash.Sum(nil))
	}
	return M
}

func print(M MerkleTree) {
	for index, val := range M.digest {
		fmt.Println(index, b64.StdEncoding.EncodeToString([]byte(val)))
	}
}

func build_tree(M MerkleTree) MerkleTree {
	hash := sha256.New()
	var i int = n - 2
	for i > 0 {
		hash.Write([]byte(M.digest[2*i+1] +
			M.digest[2*i+2]))
		M.digest[i] = string(hash.Sum(nil))
		i--

	}
	return M
}

func compressedProof(M MerkleTree, requestedLeaves []int) []string {
	marked := mark_tree(requestedLeaves)
	var i int = n - 2
	var child int
	var proof []string
	for i > 0 {
		if marked[i] {
			child = 2*i + 1
			if marked[child] {
				child += 1
			}
			if !marked[child] {
				proof = append(proof, M.digest[child])
			}
		}
		i--
	}
	return proof
}

func mark_tree(requestedLeaves []int) [twopn]bool {
	var marked [twopn]bool

	for i, _ := range requestedLeaves {
		marked[i+n-1] = true
	}
	var i int = n - 2
	for i > 0 {
		marked[i] = marked[2*i+1] || marked[2*i+2]
		i--
	}
	return marked
}

func verifyMerkle(root string, k int, s []string, indicies []int, proof []string) bool {
	hash := sha256.New()
	var child int
	var proofIndex int = 0
	marked := mark_tree(indicies)
	fmt.Println(marked)
	var tmp [twopn]string
	var defined [twopn]bool
	var i int = n - 2
	for i > 0 {
		if marked[i] {
			child = 2*i + 1
			if marked[child] {
				child += 1
			}
			if !marked[child] && proofIndex > len(proof) {

				fmt.Println("proofIndex", proofIndex)
				proofIndex += 1

				tmp[child] = proof[proofIndex]
				defined[child] = true
			}
		}
		i--
	}
	var l int = 0
	for l < k {
		tmp[indicies[l]+n] = s[i]
		defined[indicies[l]+n] = true
		l++
	}

	var j int = n - 2
	for j > 0 {
		if defined[2*i+1] && defined[2*i+2] {
			hash.Write([]byte(tmp[2*i+1] + tmp[2*i+2]))
			tmp[i] = string(hash.Sum(nil))
			defined[i] = true
		}
		j--
	}
	return defined[i] && tmp[0] == root
}
