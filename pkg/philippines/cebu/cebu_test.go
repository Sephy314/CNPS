package cebu

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCebuFlow(t *testing.T) {
	testTpks, err := Init()
	if err != nil {
		t.Errorf("Error initializing tpks: %v", err)
	}

	uid := uuid.New().String()
	ext := map[string]interface{}{
		"name": "kim",
	}

	c := Claims{
		UserId: uid,
		Ext:    ext,
	}

	cebu, err := CreateCebu(c, testTpks)

	fmt.Printf("Cebu: %+v\n", cebu)

	if err != nil {
		t.Errorf("Error creating cebu: %v", err)
	}

	// Success
	err = cebu.Verify(testTpks.KeyPair.PublicKey)
	if err != nil {
		t.Errorf("Error verifying cebu: %v", err)
	}

	// Fail
	weirdCebu := Cebu{
		Kid:         "poop",
		Philippines: "poop",
		Token:       "poop",
		ExpiresAt:   1234,
		Signature:   "poop",
		Ext:         nil,
	}

	err = weirdCebu.Verify(testTpks.KeyPair.PublicKey)
	if err == nil {
		t.Errorf("Error verifying weird cebu: %v", err)
	}

	fmt.Println("TestCebuFlow Successful! Gonna sleep rn")
}
