package crypts

import (
	"crypt/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"

	"github.com/gobuffalo/packr/v2/file/resolver/encoding/hex"
)

type TTT struct {
	aaa string
	bbb []int
}

type BBB struct {
	T *TTT
	t TTT
	b bool
}

func (bs *BBB) Get() {}
func (bs BBB) Set()  {}

var bs *BBB

func t() {

	bs = &BBB{}
	bs.Get()
}

func GenRSAKey(out io.Writer, bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateStream,
	}
	prvKeyBuf := pem.EncodeToMemory(block)
	n, err := out.Write(prvKeyBuf)
	if err != nil {
		if err == io.ErrShortWrite {
			for n < len(prvKeyBuf) {
				x, err := out.Write(prvKeyBuf[n:])
				if err != nil {
					return err
				}
				n += x
			}
		} else {
			return err
		}
		return err
	}
	return nil
}

type RsaCrypt struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// BytesToPublicKey bytes to public key
func (r *RsaCrypt) BytesToPublicKey(pub []byte) error {
	block, _ := pem.Decode(pub)
	if block == nil {
		return errors.New("private key error")
	}
	ifc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return errors.New("ifc.(*rsa.PublicKey) not ok")
	}
	r.publicKey = key
	return nil
}

// BytesToPrivateKey bytes to private key
func (r *RsaCrypt) BytesToPrivateKey(priv []byte) error {
	block, _ := pem.Decode(priv)
	if block == nil {
		return errors.New("private key error")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	r.privateKey = key
	r.publicKey = &key.PublicKey
	return nil
}
func (r *RsaCrypt) StringToPublicKey(pubStr string) error {
	return r.BytesToPublicKey([]byte(pubStr))
}
func (r *RsaCrypt) StringToPrivateKey(privStr string) error {
	return r.BytesToPrivateKey([]byte(privStr))
}

// Encrypt encrypts data with public key
func (r *RsaCrypt) Encrypt(msg []byte) ([]byte, error) {
	hash := sha512.New()
	return rsa.EncryptOAEP(hash, rand.Reader, r.publicKey, msg, nil)
}

// EncryptPKCS1v15 encrypts data with public key
func (r *RsaCrypt) EncryptPKCS1v15(msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, msg)
}

func (r *RsaCrypt) EncryptToString(plaintext string) (string, error) {
	b, err := r.Encrypt([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
func (r *RsaCrypt) EncryptToStringByHex(plaintext string) (string, error) {
	b, err := r.Encrypt([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Decrypt decrypts data with private key
func (r *RsaCrypt) Decrypt(ciphertext []byte) ([]byte, error) {
	hash := sha512.New()
	return rsa.DecryptOAEP(hash, rand.Reader, r.privateKey, ciphertext, nil)
}

// DecryptPKCS1v15 decrypts data with private key
func (r *RsaCrypt) DecryptPKCS1v15(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, ciphertext)
}

func (r *RsaCrypt) DecryptToString(ciphertext string) (string, error) {
	cipherBuf, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	b, err := r.Decrypt(cipherBuf)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *RsaCrypt) DecryptToStringByHex(ciphertext string) (string, error) {
	cipherBuf, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	b, err := r.Decrypt(cipherBuf)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
