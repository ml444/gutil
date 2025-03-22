package crypts

import (
	"encoding/base64"
	"reflect"
	"testing"

	"github.com/ml444/gutil/str"
)

func newTestAESCrypt() *AESCrypt {
	nonce := NewNonce(12)
	c, err := NewAESCrypt(
		OptWithKeyByte([]byte("u8pAdGgDU4Yw59aIFfieNiJNRrmHWYj1")),
		OptWithNonceByte(nonce),
	)
	if err != nil {
		panic(err.Error())
	}
	return c
}

func TestAESCrypt_NewNonce(t *testing.T) {
	c := newTestAESCrypt()
	nonce := c.NewNonce()
	if nonce == nil || reflect.DeepEqual(nonce, make([]byte, c.gcm.NonceSize())) {
		t.Error(" new nonce is failed")
	}
	t.Log(c.Encode2Str(nonce))
}
func TestAESCrypt_NewNonceStr(t *testing.T) {
	c := newTestAESCrypt()
	nonceStr := c.NewNonceStr()
	if nonceStr == "" || nonceStr == "AAAAAAAAAAAAAAAA" {
		t.Error(" new nonceStr is failed")
	}
	t.Log(nonceStr)
}

func TestName(t *testing.T) {
	s := "SjOYU5WCHhSeTRsFtDm65k2SRmiDxUAj"
	t.Log(s, len(s))
	b, _ := base64.StdEncoding.DecodeString(s[:len(s)-4] + "LmxQ")
	t.Log(len(s), len(b))
	ss := str.GetRandomString(32, 1)
	bb, _ := base64.StdEncoding.DecodeString(ss)
	t.Log(len(bb))
}
