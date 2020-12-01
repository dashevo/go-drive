package bls12381_test

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/bls12381"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignAndValidateBLS12381(t *testing.T) {

	privKey := bls12381.GenPrivKey()
	pubKey := privKey.PubKey()

	msg := crypto.CRandBytes(128)
	sig, err := privKey.Sign(msg)
	require.Nil(t, err)

	// Test the signature
	assert.True(t, pubKey.VerifySignature(msg, sig))
}

func TestBLSAddress(t *testing.T) {
	decodedPrivateKeyBytes, err := base64.StdEncoding.DecodeString("RokcLOxJWTyBkh5HPbdIACng/B65M8a5PYH1Nw6xn70=")
	require.Nil(t, err)
	publicKeyBytesString := "F5BjXeh0DppqaxX7a3LzoWr6CXPZcZeba6VHYdbiUCxQ23b00mFD8FRZpCz9Ug1E"
	decodedPublicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBytesString)
	require.Nil(t, err)
	decodedAddressBytes, err := hex.DecodeString("DDAD59BB10A10088C5A9CA219C3CF5BB4599B54E")
	require.Nil(t, err)
	privKey := bls12381.PrivKey(decodedPrivateKeyBytes)
	pubKey := privKey.PubKey()
	address := pubKey.Address()
	assert.EqualValues(t, decodedPublicKeyBytes, pubKey)
	assert.EqualValues(t, decodedAddressBytes, address)
}
