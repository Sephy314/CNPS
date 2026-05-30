package cebu

import (
	"crypto/ecdsa"
	"sync"
	"time"

	"github.com/google/uuid"
)

type KeyPair struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

type Philippines struct {
	Name string
	Addr string
	Tpks *Tpks
}

type Tpk struct {
	Kid       string
	Pub       string
	CreatedAt time.Time
}

type Tpks struct {
	mu sync.RWMutex

	Tpks       map[string]Tpk
	CurrentKid string
	KeyPair    KeyPair
}

func TpksInit() (*Tpks, error) {
	tpks := &Tpks{}

	t, k, e := tpks.NewTpk()
	if e != nil {
		return nil, e
	}

	tpks.Tpks = map[string]Tpk{
		t.Kid: *t,
	}

	tpks.KeyPair = *k

	return tpks, nil
}

func (t *Tpks) Rotate() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	tpk, keypair, err := t.NewTpk()
	if err != nil {
		return err
	}

	t.Tpks[tpk.Kid] = *tpk

	t.KeyPair = *keypair

	t.CurrentKid = tpk.Kid

	return nil
}

func (t *Tpks) NewTpk() (*Tpk, *KeyPair, error) {
	priv, err := GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	kid := uuid.New().String()

	pub := priv.PublicKey

	serialisedPub, err := MarshalPublicKey(&pub)

	if err != nil {
		return nil, nil, err
	}

	keyPair := KeyPair{
		PublicKey:  new(pub),
		PrivateKey: priv,
	}

	tpk := Tpk{
		Kid:       kid,
		Pub:       serialisedPub,
		CreatedAt: time.Now(),
	}

	return &tpk, &keyPair, nil

}
