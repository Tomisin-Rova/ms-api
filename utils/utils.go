package utils

import (
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

type JSON map[string]interface{}

/**
Generate Hex String, Note: Don't call this in a concurrent goroutine.
Will replace with a UUID package.
*/
func GenerateUUID(length int) (string, error) {
	// TODO: Replace this implementation with a more stable version and collision-free version.
	var src = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	// Note that err == nil only if we read len(b) bytes.
	if _, err := src.Read(b); err != nil {
		return "", err
	}

	str := hex.EncodeToString(b)
	if len(str) > length {
		return str[:length], nil
	}
	return str, nil
}

func ByteToHex(b []byte) string {
	var fish struct{}
	_ = json.Unmarshal(b, &fish)
	return hex.EncodeToString(b)
}
