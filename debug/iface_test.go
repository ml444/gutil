package debug

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestGetTypeAndDataPtr(t *testing.T) {
	type user struct {
		Name string
		Age  uint32
	}
	type args struct {
		i interface{}
	}
	x := 11
	tests := []struct {
		name  string
		args  args
		want  uintptr
		want1 interface{}
	}{
		// TODO: Add test cases.
		{"test1", args{i: &user{"foo", 21}}, 1, &user{"foo", 21}},
		{"test1", args{i: nil}, 0, nil},
		{"test1", args{i: (*int)(nil)}, 1, nil},
		{"test1", args{i: &x}, 1, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetTypeAndDataPtr(tt.args.i)
			//data := *(*int)(unsafe.Pointer(got1))
			data := (*user)(unsafe.Pointer(got1))
			t.Log(data)
			if got < tt.want {
				t.Errorf("GetTypeAndDataPtr() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(data, tt.want1) {
				t.Errorf("GetTypeAndDataPtr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
