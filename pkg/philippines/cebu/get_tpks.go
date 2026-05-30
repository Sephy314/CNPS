package cebu

import (
	"crypto/ecdsa"

	"github.com/Sephy314/cnps/pkg/philippines/errs"
)

func (t *Tpks) GetCurrentPrivateKey() (*ecdsa.PrivateKey, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	priv := t.KeyPair.PrivateKey

	return priv, nil
}

func (t *Tpks) GetPublicKeyByKid(kid string) (*ecdsa.PublicKey, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	pub, ok := t.Tpks[kid]
	if !ok {
		return nil, errs.NotFoundError
	}

	unmarshalled, err := UnmarshalPublicKey(pub.Pub)

	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}
