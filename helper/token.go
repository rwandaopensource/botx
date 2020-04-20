package helper

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"
)

var (
	// PrivateKey is the private key
	PrivateKey ed25519.PrivateKey = MustDecodeKey(os.Getenv("PRIVATE_KEY"))
	// PublicKey is the private key
	PublicKey ed25519.PublicKey = MustDecodeKey(os.Getenv("PUBLIC_KEY"))
)

// GenerateKey generate new private and public key
func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(nil)
}

// Sign signs the message with privateKey and returns a signature.
// it returns nil of []byte and error if the size of the private key is not equal to 64
func Sign(priv ed25519.PrivateKey, message []byte) ([]byte, error) {
	if len(priv) != ed25519.PrivateKeySize {
		return nil, errors.New("Private key is invalid")
	}
	return ed25519.Sign(priv, message), nil
}

// Verify reports whether sig is a valid signature of message by publicKey.
// it returns nil of false and error if the size of the public key is not equal to 32
func Verify(pub ed25519.PublicKey, message, sign []byte) (bool, error) {
	if len(pub) != ed25519.PublicKeySize {
		return false, errors.New("Public key or signature are invalid")
	}
	return ed25519.Verify(pub, message, sign), nil
}

// EncodeKey from byte to be a decent readable string
func EncodeKey(key []byte) string {
	return hex.EncodeToString(key)
}

// DecodeKey decode string key into bytes
func DecodeKey(key string) (d []byte, err error) {
	d, err = hex.DecodeString(key)
	return
}

// MustDecodeKey decode key or panic
func MustDecodeKey(key string) []byte {
	d, err := hex.DecodeString(key)
	PanicError(err, "")
	return d
}

// PrintKey it prints generated private and public key or error
func PrintKey() {
	pub, priv, err := GenerateKey()
	PanicError(err, "")
	Print("private key: ", EncodeKey(pub))
	Print("public key: ", EncodeKey(priv))
}

// ClientIDAndSecretKey generate new signed client and secret key
func ClientIDAndSecretKey(prv ed25519.PrivateKey) (ID string, secret []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		return
	}()
	ID = strings.Replace(uuid.New().String(), "-", "", -1)
	secret, err = Sign(prv, []byte(ID))
	return
}
