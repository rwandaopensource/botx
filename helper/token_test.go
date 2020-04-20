package helper_test

import (
	"bytes"
	"crypto/ed25519"
	"testing"

	"github.com/rwandaopensource/botx/helper"
)

func TestED25519PublicAndPrivate(t *testing.T) {
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
}

func TestED25519EncodeDecode(t *testing.T) {
	pub, priv, err := helper.GenerateKey()
	if err != nil {
		helper.TestError(t, err)
	}
	token, err := helper.Sign(priv, []byte("hello"))
	if err != nil {
		helper.TestError(t, err)
	}
	r, err := helper.Verify(pub, []byte("hello"), token)
	if !r || err != nil {
		helper.TestError(t, err)
	}
}

func TestClientAndSecretkey(t *testing.T) {
	pub, priv, err := helper.GenerateKey()
	clientID, clientSecret, err := helper.ClientIDAndSecretKey(priv)
	if err != nil {
		helper.TestError(t, err)
	}
	if v, err := helper.Verify(pub, []byte(clientID), clientSecret); !v || err != nil {
		if err == nil {
			helper.TestError(t, "invalid signature")
		} else {
			helper.TestError(t, err)
		}
	}
}
