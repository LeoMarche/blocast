package exchange

import "crypto/rsa"

//Transaction contains the informations relative to
//a transaction
type Transaction struct {
	PublicKey *rsa.PublicKey
	Object    string
	Amount    int
}

//SignedTransaction holds a transaction and its signature
type SignedTransaction struct {
	T   Transaction
	sig []byte
}

func Exchange() {
	return
}
