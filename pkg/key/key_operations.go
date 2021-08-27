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
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	//Hash Digest
	h := sha512.New()
	_, err = h.Write(b.Bytes())
	if err != nil {
		return nil, err
	}
	hs := h.Sum(nil)

	//Signature
	return rsa.SignPSS(rand.Reader, privKey, crypto.SHA512, hs, nil)
}

func VerifyStruct(v interface{}, pubKey *rsa.PublicKey, signature []byte) (bool, error) {
	//Serialization
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return false, err
	}

	//Hash Digest
	h := sha512.New()
	_, err = h.Write(b.Bytes())
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
