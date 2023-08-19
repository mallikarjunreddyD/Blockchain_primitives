package main

import "fmt"

func genTransactions(userData map[int]user) {
	/** Genesis Transactions */
	execute("100", userData[0].Address) //(100,Alice)
	execute("50", userData[0].Address)  //(50,Alice)
	execute("70", userData[1].Address)  //(70,Bob)
	execute("60", userData[2].Address)  //(60,Charlie)

	fmt.Println("UTXOList after Genesis")
	printList(UTXOList)

	/** Transaction 1 */
	input1Id := "0x73b45126" //(100,Alice)
	input1Sig := genSignature(userData[0].PrivateKey, input1Id)
	input2Id := "0xd49e7d8b" //(50,Alice)
	input2Sig := genSignature(userData[0].PrivateKey, input2Id)

	Trans1Inputs := []input{{
		Id:        input1Id,
		signature: input1Sig,
	}, {
		Id:        input2Id,
		signature: input2Sig,
	}}

	Trans1Outputs := []UTXO{{ //((120,Darth),(30,Alice))
		value: "120",
		owner: userData[3].Address,
	}, {
		value: "30",
		owner: userData[0].Address,
	}}

	/** Transaction 2*/
	input1Id = "0x6971747c" //(120,Darth)
	input1Sig = genSignature(userData[3].PrivateKey, input1Id)
	Trans2Inputs := []input{{
		Id:        input1Id,
		signature: input1Sig,
	}}
	Trans2Outputs := []UTXO{{ //((50,Earl),(30,Bob),(30,Fin))
		value: "50",
		owner: userData[4].Address,
	}, {
		value: "30",
		owner: userData[1].Address,
	}, {
		value: "30",
		owner: userData[5].Address,
	}}

	/** Transaction 3 */
	input1Id = "0x737f578d" //(30,Bob)
	input1Sig = genSignature(userData[1].PrivateKey, input1Id)
	input2Id = "0x1fad6b7c" //(70,Bob)
	input2Sig = genSignature(userData[1].PrivateKey, input2Id)

	Trans3Inputs := []input{{
		Id:        input1Id,
		signature: input1Sig,
	}, {
		Id:        input2Id,
		signature: input2Sig,
	}}

	Trans3Outputs := []UTXO{{ //((60,Grace),(60,Alice))
		value: "60",
		owner: userData[3].Address,
	}, {
		value: "60",
		owner: userData[0].Address,
	}}

	/** Transaction 1 execution */
	executeTransaction(Trans1Inputs, Trans1Outputs)
	fmt.Println("UTXOList after Transaction 1")
	printList(UTXOList)

	/** Transaction 2 execution */
	executeTransaction(Trans2Inputs, Trans2Outputs)
	fmt.Println("UTXOList after Transaction 2")
	printList(UTXOList)

	/** Transaction 3 execution */
	executeTransaction(Trans3Inputs, Trans3Outputs)
	fmt.Println("UTXOList after Transaction 3")
	printList(UTXOList)

}
