package crypts

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(dst string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(dst))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256WithSalt(dst string, salt string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(dst + salt))
	return hex.EncodeToString(h.Sum(nil))
}
