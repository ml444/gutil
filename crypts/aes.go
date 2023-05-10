package crypts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

type stringer interface {
	EncodeToString([]byte) string
	DecodeString(string) ([]byte, error)
}

type OptFunc func(c *AESCrypt)

func OptWithKeyByte(key []byte) OptFunc {
	return func(c *AESCrypt) {
		c.key = key
	}
}

func OptWithNonceByte(nonce []byte) OptFunc {
	return func(c *AESCrypt) {
		c.nonce = nonce
	}
}
func OptWithDataByte(data []byte) OptFunc {
	return func(c *AESCrypt) {
		c.data = data
	}
}

func OptWithIStr(istr stringer) OptFunc {
	return func(c *AESCrypt) {
		c.istr = istr
	}
}

type AESCrypt struct {
	key   []byte
	nonce []byte
	data  []byte
	gcm   cipher.AEAD
	istr  stringer
}

func NewAESCrypt(opts ...OptFunc) (*AESCrypt, error) {
	c := &AESCrypt{}
	for _, opt := range opts {
		opt(c)
	}
	if c.istr == nil {
		c.istr = base64.RawStdEncoding
	}
	err := c.initGCM()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *AESCrypt) NewNonce() []byte {
	return NewNonce(c.gcm.NonceSize())
}

func (c *AESCrypt) NewNonceStr() string {
	return c.istr.EncodeToString(NewNonce(c.gcm.NonceSize()))
}

func (c *AESCrypt) Encode2Str(b []byte) string {
	return c.istr.EncodeToString(b)
}

func (c *AESCrypt) Decode2Byte(s string) ([]byte, error) {
	return c.istr.DecodeString(s)
}

func (c *AESCrypt) initGCM() error {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return err
	}
	c.gcm, err = cipher.NewGCM(block)
	if err != nil {
		return err
	}
	return nil
}

func (c *AESCrypt) Encrypt(plaintext string) (string, error) {
	return c.Encrypt4Byte([]byte(plaintext))
}

func (c *AESCrypt) Encrypt4Byte(plaintext []byte) (string, error) {
	if c.gcm == nil {
		err := c.initGCM()
		if err != nil {
			return "", err
		}
	}
	ciphertext := c.gcm.Seal(nil, c.nonce, plaintext, c.data)
	return c.istr.EncodeToString(ciphertext), nil
}

func (c *AESCrypt) Decrypt(ciphertext string) (string, error) {
	plaintext, err := c.Decrypt2Byte(ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func (c *AESCrypt) Decrypt2Byte(ciphertext string) ([]byte, error) {
	b, err := c.istr.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	plaintext, err := c.gcm.Open(nil, c.nonce, b, c.data)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func NewNonce(size int) []byte {
	nonce := make([]byte, size)
	_, _ = rand.Read(nonce)
	return nonce
}
