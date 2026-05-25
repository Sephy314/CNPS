package cebu

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}
