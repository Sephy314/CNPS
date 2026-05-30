package cebu

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"time"

	"github.com/Sephy314/cnps/pkg/philippines/errs"
)

type Cebu struct {
	Kid         string         `json:"kid"`
	Philippines string         `json:"philippines"`
	Token       string         `json:"token"`
	ExpiresAt   int64          `json:"expires_at"`
	Signature   string         `json:"signature"`
	Ext         map[string]any `json:"ext,omitempty"`
}

type Claims struct {
	UserId string         `json:"user_id"`
	Ext    map[string]any `json:"ext,omitempty"`
}

type ECDSASignature struct {
	R, S *big.Int
}

func getExpireAt(t time.Time) time.Time {
	exp := time.Minute * 5

	return t.Add(exp)
}

func (c *Cebu) Sign(priv *ecdsa.PrivateKey) error {
	// deterministic serialization
	data, err := json.Marshal(struct {
		Kid         string `json:"kid"`
		Philippines string `json:"philippines"`
		Token       string `json:"token"`
	}{
		Kid:         c.Kid,
		Philippines: c.Philippines,
		Token:       c.Token,
	})
	if err != nil {
		return err
	}

	// hash
	hash := sha256.Sum256(data)

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return err
	}

	// ASN.1 encode
	sig, err := asn1.Marshal(ECDSASignature{
		R: r,
		S: s,
	})
	if err != nil {
		return err
	}

	c.Signature = base64.StdEncoding.EncodeToString(sig)

	return nil
}

func (c *Cebu) Verify(pub *ecdsa.PublicKey) error {
	// same deterministic serialization
	data, err := json.Marshal(struct {
		Kid         string `json:"kid"`
		Philippines string `json:"philippines"`
		Token       string `json:"token"`
	}{
		Kid:         c.Kid,
		Philippines: c.Philippines,
		Token:       c.Token,
	})
	if err != nil {
		return err
	}

	// hash
	hash := sha256.Sum256(data)

	// decode signature
	sigBytes, err := base64.StdEncoding.DecodeString(c.Signature)
	if err != nil {
		return err
	}

	var sig ECDSASignature

	_, err = asn1.Unmarshal(sigBytes, &sig)
	if err != nil {
		return err
	}

	// verify
	valid := ecdsa.Verify(pub, hash[:], sig.R, sig.S)

	if !valid {
		return errs.InvalidSignatureError
	}

	// verify expired
	if c.ExpiresAt < time.Now().Unix() {
		return errs.ExpiredSignatureError
	}

	return nil
}
func CreateCebu(
	server Philippines,
	c Claims,
	tpks *Tpks,
) (*Cebu, error) {

	// serialize claims
	claimsBytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	// create token
	token := base64.RawURLEncoding.EncodeToString(claimsBytes)

	cebu := &Cebu{
		Kid:         tpks.CurrentKid,
		Philippines: server.Name,

		Token: token,

		ExpiresAt: getExpireAt(time.Now()).Unix(),

		Ext: c.Ext,
	}

	// sign
	err = cebu.Sign(server.Tpks.KeyPair.PrivateKey)
	if err != nil {
		return nil, err
	}

	return cebu, nil
}

func (c *Cebu) GetUID() (string, error) {
	// decode token
	tokenBytes, err := base64.RawURLEncoding.DecodeString(c.Token)
	if err != nil {
		return "", err
	}

	// parse claims
	var claims Claims

	err = json.Unmarshal(tokenBytes, &claims)
	if err != nil {
		return "", err
	}

	return claims.UserId, nil
}
