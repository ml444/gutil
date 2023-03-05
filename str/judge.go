package str

import (
	"unicode"
)

func IsAnyEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}
	return false
}

func IsAllEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) > 0 {
			return false
		}
	}
	return true
}

// IsSameByRegroup  Whether the two strings are the same after sorting and reorganization
func IsSameByRegroup(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var totalNum1 int
	for _, v := range s1 {
		totalNum1 += int(v)
	}
	var totalNum2 int
	for _, v := range s2 {
		totalNum2 += int(v)
	}
	if totalNum1 != totalNum2 {
		return false
	}
	return true
}

func IsAllLetter(s string) bool {
	for _, v := range s {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

func IsCapitalizedFirst(s string) bool {
	if d := s[0]; d >= 'A' && d <= 'Z' {
		return true
	}
	return false
}
