package str

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	RandomStringModNumberPlusLetter           = 1
	RandomStringModNumberPlusLetterPlusSymbol = 2
	RandomStringModNumber                     = 3
)

func GenRandomStr() string {
	rndStr := fmt.Sprint(
		os.Getpid(), time.Now().UnixNano(), rand.Float64())
	h := md5.New()
	_, _ = io.WriteString(h, rndStr)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetRandomString(length int64, mod uint32) string {
	var strKey string
	if mod == RandomStringModNumberPlusLetter {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if mod == RandomStringModNumberPlusLetterPlusSymbol {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*(){}|\\[]?/"
	} else if mod == RandomStringModNumber {
		strKey = "0123456789"
	} else {
		strKey = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = strKey[r.Intn(len(strKey))]
	}

	return string(bytes)
}

func GenCode(n int) string {
	rand.Seed(time.Now().UnixNano())
	l := make([]rune, n+1)
	l[0] = '1'
	for i := 0; i < n; i++ {
		l[i+1] = '0'
	}
	i64, _ := strconv.ParseInt(string(l), 10, 64)
	return fmt.Sprintf("%0"+fmt.Sprintf("%dv", n), rand.Int63n(i64))
	//return fmt.Sprintf("%0"+fmt.Sprintf("%dv", n), rand.Intn(10*n))
}
