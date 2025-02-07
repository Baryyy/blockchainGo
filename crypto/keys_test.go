package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("True")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	// Invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("false")))

	//Invalid pubKey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))

}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "f6c5fbc8111c3c0c2583b0a52d6cb6b00baa93d9a9a3e56aa914c50ad23fa31c"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "daf866c2551ac4b0a1f569d50dac00e62934be4a"
	)
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())

}
