package tpksStore

import "time"

type TokenStore interface {
	Save(token string, value []byte, ttl time.Duration) error
	Get(token string) ([]byte, error)
	Delete(token string) error
}
