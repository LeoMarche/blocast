package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
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
	publicKeyBytes := x509.MarshalPKCS1PublicKey(pubKey)
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

//RetrieveKeys is a function that retrieves RSA Key Pair with pem files
func RetrieveKeys(privKeyPath, pubKeyPath string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	//Opening PEM files
	privKeyBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, nil, err
	}

	//Decoding PEM chains
	privKeyPEM, _ := pem.Decode(privKeyBytes)
	pubKeyPEM, _ := pem.Decode(pubKeyBytes)

	//Parsing RSA keys
	privateKey, err := x509.ParsePKCS1PrivateKey(privKeyPEM.Bytes)
	if err != nil {
		return nil, nil, err
	}
	publicKey, err := x509.ParsePKCS1PublicKey(pubKeyPEM.Bytes)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}
