package str

import (
	"bytes"
	"strings"
)

// CamelToSnake 驼峰转蛇形
func CamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// SnakeToCamel 蛇形转驼峰
func SnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func CutStr(s string, max int) string {
	sr := []rune(s)
	if len(sr) > max {
		sr = sr[:max]
	}
	return string(sr)
}

// MixStrEncode confuse string to encrypt them
func MixStrEncode(s string) string {
	var isDouble bool
	b := []byte(s)
	bLen := len(b)
	rang := bLen / 2
	if bLen%2 == 0 {
		isDouble = true
	}
	var buffer bytes.Buffer
	for i := 0; i < rang; i++ {
		j := bLen - 1 - i
		buffer.Write(b[i : i+1])
		if isDouble || j != rang {
			buffer.Write(b[j : j+1])
		}
	}
	return buffer.String()
}

// MixStrDecode decrypting confused string
func MixStrDecode(s string) string {
	b := []byte(s)
	bLen := len(b)
	var buffer bytes.Buffer
	for i := 0; i < bLen; i = i + 2 {
		buffer.Write(b[i : i+1])
	}
	for i := bLen - 1; i >= 0; i = i - 2 {
		buffer.Write(b[i : i+1])
	}
	return buffer.String()
}

// ToUpperFirst Capitalize the first letter of a string
func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}
	if d := s[0]; d >= 'a' && d <= 'z' {
		return string(d-32) + s[1:]
	}
	return s
}

// ToLowerFirst Make the first letter of a string lowercase
func ToLowerFirst(s string) string {
	if s == "" {
		return s
	}
	if d := s[0]; d >= 'A' && d <= 'Z' {
		return string(d+32) + s[1:]
	}
	return s
}

// InvertStr A string flips the string around the middle character axis
func InvertStr(s string) string {
	runeList := []rune(s)
	rLen := len(runeList)
	for i := 0; i < rLen/2; i++ {
		runeList[i], runeList[rLen-1-i] = runeList[rLen-1-i], runeList[i]
	}
	return string(runeList)
}

func ReplaceBlank(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
