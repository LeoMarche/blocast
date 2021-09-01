package key

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_DIR = "test"

func TestStoreRetrieve(t *testing.T) {
	os.MkdirAll(TEST_DIR, os.ModePerm)
	defer os.RemoveAll(TEST_DIR)
	privFile := path.Join(TEST_DIR, "priv.pem")
	pubFile := path.Join(TEST_DIR, "pub.pem")
	priv, pub, _ := GenerateKeys(2048)

	StoreKeys(priv, pub, privFile, pubFile)
	priv2, pub2, _ := RetrieveKeys(privFile, pubFile)

	assert.Equalf(t, priv, priv2, "Bad private key : %v instead of %v", priv2, priv)
	assert.Equalf(t, pub, pub2, "Bad public key : %v instead of %v", pub2, pub)
}

func TestSignVerify(t *testing.T) {
	type t1 struct {
		I int
		S string
	}
	type t2 struct {
		I *int
		S *string
	}
	type test struct {
		V interface{}
		C interface{}
		R bool
		E bool
	}

	privKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	test1 := test{"sametext", "sametext", true, false}

	var tests = []test{test1}

	for _, v := range tests {
		sig, _ := SignStruct(v.V, privKey)
		res, err := VerifyStruct(v.C, &privKey.PublicKey, sig)
		assert.Equalf(t, v.R, res, "Bad resuls returned for test %v", v.V)
		assert.Equalf(t, v.E, err != nil, "Bad error for test %v", v.V)
	}
}
