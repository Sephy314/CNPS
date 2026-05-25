package cebu

import (
	"crypto/ecdsa"

	"github.com/Sephy314/cnps/pkg/philippines/errs"
)

func (tpks *Tpks) GetCurrentPrivateKey() (*ecdsa.PrivateKey, error) {
	tpks.mu.Lock()
	defer tpks.mu.Unlock()

	priv := tpks.KeyPair.PrivateKey

	return priv, nil
}

func (tpks *Tpks) GetPublicKeyByKid(kid string) (*ecdsa.PublicKey, error) {
	tpks.mu.Lock()
	defer tpks.mu.Unlock()

	pub := tpks.Tpks[kid].Pub
	if pub == nil {
		return nil, errs.NotFoundError
	}

	return pub, nil
}
