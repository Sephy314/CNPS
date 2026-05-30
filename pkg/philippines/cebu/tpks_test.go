package cebu

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTpksFlow(t *testing.T) {
	tpks, err := TpksInit()
	if err != nil {
		t.Fatalf("Error initializing TPKS: %v", err)
	}

	old_kid := tpks.CurrentKid

	err = tpks.Rotate()
	if err != nil {
		t.Fatalf("Rotate failed: %v", err)
	}

	if old_kid == tpks.CurrentKid {
		t.Fatalf("Kid not rotated. old : %v, new : %v ", old_kid, tpks.CurrentKid)
	}

	// When it finds exist one
	pub, err := tpks.GetPublicKeyByKid(tpks.CurrentKid)
	if err != nil {
		t.Fatalf("Error getting public key : %v", err)
	}

	if !reflect.DeepEqual(pub, tpks.KeyPair.PublicKey) {
		t.Fatalf("Wrong public key. expected : %v, got : %v", tpks.KeyPair.PublicKey, pub)
	}

	errKid := "WEIRD PUB KEY"

	errPub, err := tpks.GetPublicKeyByKid(errKid)

	if errPub != nil && err == nil {
		t.Fatalf("Public key shouldn't be found: %v", errPub)
	}

	fmt.Println("Hooray! Test Successful! Gonna have some fried chicken!")

}
