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

type Tpk struct {
	Kid       string
	Pub       *ecdsa.PublicKey
	CreatedAt time.Time
}

type Tpks struct {
	mu sync.RWMutex

	Tpks       map[string]Tpk
	CurrentKid string
	KeyPair    KeyPair
}

func Init() (*Tpks, error) {
	priv, err := GenerateKey()
	if err != nil {
		return nil, err
	}

	kid := uuid.New().String()

	pub := priv.PublicKey

	keyPair := KeyPair{
		PublicKey:  new(pub),
		PrivateKey: priv,
	}

	tpk := Tpk{
		Kid:       kid,
		Pub:       &pub,
		CreatedAt: time.Now(),
	}

	tpks := Tpks{
		CurrentKid: kid,
		KeyPair:    keyPair,
	}

	tpks.Tpks = map[string]Tpk{
		kid: tpk,
	}

	return &tpks, nil
}

func (tpks *Tpks) Rotate() error {
	tpks.mu.Lock()
	defer tpks.mu.Unlock()

	priv, err := GenerateKey()
	if err != nil {
		return err
	}

	kid := uuid.New().String()

	pub := priv.PublicKey

	keyPair := KeyPair{
		PublicKey:  new(pub),
		PrivateKey: priv,
	}

	tpk := Tpk{
		Kid:       kid,
		Pub:       &pub,
		CreatedAt: time.Now(),
	}

	tpks.Tpks[kid] = tpk

	tpks.KeyPair = keyPair

	tpks.CurrentKid = kid

	return nil
}
