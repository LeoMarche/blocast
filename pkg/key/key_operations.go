//key is the package that manages the rsa keys needed to sign all
//the operations
package key

import (
	"crypto/rand"
	"crypto/rsa"
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
