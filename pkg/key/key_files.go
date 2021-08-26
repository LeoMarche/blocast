package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//StoreKeys is a function that stores a rsa key pair to pem files
func StoreKeys(privKey *rsa.PrivateKey, pubKey *rsa.PublicKey, privKeyPath, pubKeyPath string) error {
	//Creating private key pem
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create(privKeyPath)
	if err != nil {
		return err
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		return err
	}

	//Creating public key pem
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create(pubKeyPath)
	if err != nil {
		return err
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		return err
	}

	return nil
}
