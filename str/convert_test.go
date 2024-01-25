package str

import (
	"testing"
)

func TestName(t *testing.T) {
	//s := "AccountServiceName"
	s := "account_service_name"
	t.Log(ToUpperFirst(s))
	t.Log(ToUpperFirst(s))
	t.Log('A')
	t.Log('a')
}

func BenchmarkCamelToSnake(b *testing.B) {
	s := "account_service_name"
	for i := 0; i < b.N; i++ {
		ToUpperFirst(s)
	}
}

//func BenchmarkCamelToSnake2(b *testing.B) {
//	s := "account_service_name"
//	for i := 0; i < b.N; i++ {
//		ToUpperFirst2(s)
//	}
//}

func TestMixStrEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test1", "abcdefg1234567", "a7b6c5d4e3f2g1"},
		{"test2", "abcdefg12345678", "a8b7c6d5e4f3g21"},
		{"test3", "abcdefgH1234567", "a7b6c5d4e3f2g1H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MixStrEncode(tt.s); got != tt.want {
				t.Errorf("MixStrEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpperFirst(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"tst1", "abc", "Abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpperFirst(tt.s); got != tt.want {
				t.Errorf("ToUpperFirst2() = %v, want %v", got, tt.want)
			}
		})
	}
}
