package helper_test

import (
	"bytes"
	"crypto/ed25519"
	"testing"

	"github.com/rwandaopensource/botx/pkg/helper"
)

func TestED25519EncodeDecode(t *testing.T) {
	var pub ed25519.PublicKey = []byte{}
	var priv ed25519.PrivateKey = []byte{}
	var err error
	pub, priv, err = helper.GenerateKey()
	if err != nil {
		helper.TestErrorf(t, "expected %v found %s", nil, err.Error())
	}
	if !bytes.Equal([]byte(priv.Public().(ed25519.PublicKey)), []byte(pub)) {
		helper.TestError(t, "expected newly generated private and public key to match")
	}
	if !bytes.Equal([]byte(priv.Public().(ed25519.PublicKey)), []byte(pub)) {
		helper.TestError(t, "expected app's private and public key to match")
	}
}
