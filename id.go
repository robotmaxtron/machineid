package machineid

import (
	"crypto/hmac"
	"encoding/hex"
	"fmt"
	"hash"
	"sync"

	"golang.org/x/crypto/sha3"
)

// ID returns the platform-specific machine id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID instead.
func ID() (string, error) {
	id, err := getUID()
	if err != nil {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return id, nil
}

// ProtectedID returns a stable, printable identifier derived from the machine ID.
// It computes HMAC-SHA3-512(appID) keyed by the machine ID and hex-encodes the result.
func ProtectedID(appID string) (string, error) {
	id, err := getUID()
	if err != nil {
		return "", err
	}

	mac := hmac.New(func() hash.Hash {
		h := sha3.New512()
		return h
	}, []byte(id))

	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func getUID() (string, error) {
	var (
		once      sync.Once
		cachedID  string
		cachedErr error
	)
	once.Do(func() {
		uid, err := machineID()
		if err != nil {
			cachedErr = err
			return
		}
		cachedID = uid
	})
	if cachedErr != nil {
		return "", cachedErr
	}
	return cachedID, nil
}
