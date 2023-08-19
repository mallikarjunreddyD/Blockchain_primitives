package main

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

type input struct {
	Id        string
	signature string
}
