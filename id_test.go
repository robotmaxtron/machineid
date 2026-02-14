package machineid

import (
	"encoding/hex"
	"testing"
)

func Test_ID(t *testing.T) {
	got, err := ID()
	if err != nil {
		t.Error(err)
	}
	if got == "" {
		t.Error("got empty machine id")
	}
}

func Test_ProtectedID_format(t *testing.T) {
	got, err := ProtectedID("ms.azur.appX")
	if err != nil {
		t.Error(err)
	}
	if got == "" {
		t.Error("protected id should not be empty")
	}

	// BLAKE2b-512 output is 64 bytes => 128 hex chars.
	if len(got) != 128 {
		t.Fatalf("protected id length = %d, want %d", len(got), 128)
	}
	if _, err := hex.DecodeString(got); err != nil {
		t.Fatalf("protected id is not valid hex: %v", err)
	}
}

func Test_ProtectedID_changes_with_appID(t *testing.T) {
	a, err := ProtectedID("appA")
	if err != nil {
		t.Error(err)
	}
	b, err := ProtectedID("appB")
	if err != nil {
		t.Error(err)
	}
	if a == b {
		t.Error("expected different protected ids for different appIDs")
	}
}
