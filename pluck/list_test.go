package pluck

import (
	"reflect"
	"testing"
)

type testUser struct {
	Id   uint64
	Name string
	Age  uint
}

func TestToNumbers(t *testing.T) {
	type args struct {
		list      any
		fieldName string
		unique    bool
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				list:      []*testUser{{1, "foo", 12}, {2, "bar", 18}, {3, "abc", 18}},
				fieldName: "Age",
				unique:    false,
			},
			want: []uint{12, 18, 18},
		},
		{
			name: "",
			args: args{
				list:      []*testUser{{1, "foo", 12}, {2, "bar", 18}, {3, "abc", 18}},
				fieldName: "Age",
				unique:    true,
			},
			want: []uint{12, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNumbers[uint](tt.args.list, tt.args.fieldName, tt.args.unique); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
