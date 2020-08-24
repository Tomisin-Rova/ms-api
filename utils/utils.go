package utils

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

func Pack(in interface{}, target interface{}) error {
	var e1 error
	var b []byte
	switch in.(type) {
	case []byte:
		b = in.([]byte)
	// Do something.
	default:
		// Do the rest.
		b, e1 = json.Marshal(in)
		if e1 != nil {
			return e1
		}
	}

	buf := bytes.NewBuffer(b)
	enc := json.NewDecoder(buf)
	enc.UseNumber()
	if err := enc.Decode(&target); err != nil {
		return err
	}
	return nil
}

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
