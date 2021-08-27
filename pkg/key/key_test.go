package key

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
