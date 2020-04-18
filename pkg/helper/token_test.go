package helper_test

import (
	"bytes"
	"testing"

	"github.com/rwandaopensource/botx/pkg/helper"
)

func TestED25519EncodeDecode(t *testing.T) {
	pub, priv, err := helper.GenerateKey()
	if err != nil {
		t.Errorf("expected %v found %ds", nil, err.Error())
	}
	if bytes.Equal(priv.Public()[:], pub) {
		t.Error("expected newly generated private and publuc key to match")
	}
	if bytes.Equal(helper.PrivateKey.Public()[:], helper.PublicKey) {
		t.Error("expected app's private and public key to match")
	}
}
