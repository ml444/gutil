package str

import (
	"testing"
)

func TestGenCode(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenCode(10))
	}
}
