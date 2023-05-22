package str

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash/fnv"
)

func HashCode(s string) uint32 {
	f := fnv.New32a()
	_, _ = f.Write([]byte(s))
	return f.Sum32()
}

func Md5(s string) string {
	m := md5.New()
	_, _ = m.Write([]byte(s))
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Sha256(dst string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(dst))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256WithSalt(dst string, salt string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(dst + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}
