package str

import "testing"

func TestIsSameByRegroup(t *testing.T) {
	s1 := "abcdefghijklmn opqrstuvwxyz"
	s2 := "qwertyuiop asdfghjklzxcvbnm"
	result := IsSameByRegroup(s1, s2)
	if result != true {
		t.Fatal("failed")
	}
}

//func BenchmarkIsSameByRegroup(b *testing.B) {
//	s1 := "abcdefghijklmn opqrstuvwxyz"
//	s2 := "qwertyuiop asdfghjklzxcvbnm"
//	for i := 0; i < b.N; i++ {
//		_ = IsSameByRegroup(s1, s2)
//	}
//}
//func BenchmarkIsSameByRegroup2(b *testing.B) {
//	s1 := "abcdefghijklmn opqrstuvwxyz"
//	s2 := "qwertyuiop asdfghjklzxcvbnm"
//	for i := 0; i < b.N; i++ {
//		_ = IsSameByRegroup(s1, s2)
//	}
//}
