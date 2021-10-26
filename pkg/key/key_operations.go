//key is the package that manages the rsa keys needed to sign all
//the operations
package key

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/gob"
)

func serializeStruct(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

//GenerateKeys is a function used to generate rsa private and public keys
//with a given size
func GenerateKeys(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	return privKey, &privKey.PublicKey, nil
}

//SignStruct serializes a struct using gob and signs it using the given private key
func SignStruct(v interface{}, privKey *rsa.PrivateKey) ([]byte, error) {

	//Serialization
	b, err := serializeStruct(v)
	if err != nil {
		return nil, err
	}

	//Hash Digest
	h := sha512.New()
	_, err = h.Write(b)
	if err != nil {
		return nil, err
	}
	hs := h.Sum(nil)

	//Signature
	return rsa.SignPSS(rand.Reader, privKey, crypto.SHA512, hs, nil)
}

func VerifyStruct(v interface{}, pubKey *rsa.PublicKey, signature []byte) (bool, error) {

	//Serialization
	b, err := serializeStruct(v)
	if err != nil {
		return false, err
	}

	//Hash Digest
	h := sha512.New()
	_, err = h.Write(b)
	if err != nil {
		return false, err
	}
	hs := h.Sum(nil)

	//Verification
	err = rsa.VerifyPSS(pubKey, crypto.SHA512, hs, signature, nil)
	if err != nil {
		return false, nil
	}
	return true, nil
}

//Transaction contains the informations relative to
//a transaction
type Transaction struct {
	PublicKey *rsa.PublicKey
	Object    string
	Amount    int
	ID        int
}

func (t *Transaction) Sign(privKey *rsa.PrivateKey) (*SignedTransaction, error) {
	sig, err := SignStruct(t, privKey)
	if err != nil {
		return nil, err
	}
	st := SignedTransaction{
		T:   *t,
		sig: sig,
	}
	return &st, nil
}

//SignedTransaction holds a transaction and its signature
type SignedTransaction struct {
	T   Transaction
	sig []byte
}

func (st *SignedTransaction) Serialize() ([]byte, error) {
	return serializeStruct(st)
}
