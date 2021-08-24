package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

var KEY_FOLDER string = "keys"
var INIT_FOLDERS = []string{KEY_FOLDER}
var KEY_SIZE int = 8192
var KEY_NAME string = "block-cast"

func generateKeys(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	return privKey, &privKey.PublicKey, nil
}

func storeKeys(privKey *rsa.PrivateKey, pubKey *rsa.PublicKey, privKeyPath, pubKeyPath string) error {
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

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func initializeFolder(folderPath string) error {
	ex, err := exists(folderPath)
	if err != nil {
		return err
	}

	if !ex {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func initializeFolders(folderPaths []string) error {
	for _, s := range folderPaths {
		err := initializeFolder(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := initializeFolders(INIT_FOLDERS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privKey, pubKey, err := generateKeys(KEY_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	basename := filepath.Join(KEY_FOLDER, KEY_NAME)
	err = storeKeys(privKey, pubKey, basename+"-private.pem", basename+"-public.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
