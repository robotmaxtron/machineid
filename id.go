package machineid

import (
	"crypto/hmac"
	"encoding/hex"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2b"
)

// ID returns the platform-specific machine id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID instead.
func ID() (string, error) {
	id, err := machineID()
	if err != nil {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return id, nil
}

// ProtectedID returns a stable, printable identifier derived from the machine ID.
// It computes HMAC-BLAKE2b-512(appID) keyed by the machine ID and hex-encodes the result.
func ProtectedID(appID string) (string, error) {
	id, err := ID()
	if err != nil {
		return "", err
	}

	mac := hmac.New(func() hash.Hash {
		// New512 never returns nil; it returns an error only for invalid params (none here).
		h, _ := blake2b.New512(nil)
		return h
	}, []byte(id))

	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil)), nil
}
