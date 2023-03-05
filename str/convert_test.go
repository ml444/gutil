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
func BenchmarkCamelToSnake2(b *testing.B) {
	s := "account_service_name"
	for i := 0; i < b.N; i++ {
		ToUpperFirst(s)
	}
}
