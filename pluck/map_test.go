package pluck

import (
	"reflect"
	"testing"
)

var testUserList = []*testUser{
	{Id: 1, Name: "user1", Age: 18},
	{Id: 2, Name: "user2", Age: 18},
	{Id: 3, Name: "user3", Age: 18},
	{Id: 4, Name: "user4", Age: 18},
	{Id: 5, Name: "user5", Age: 25},
	{Id: 6, Name: "user6", Age: 25},
	{Id: 7, Name: "user7", Age: 30},
	{Id: 8, Name: "user8", Age: 30},
	{Id: 9, Name: "user9", Age: 18},
}

func TestStructToMap(t *testing.T) {
	type args struct {
		s          any
		removeZero bool
		deep       bool
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "",
			args: args{
				s:          &testUser{1, "foo", 18},
				removeZero: false,
				deep:       false,
			},
			want: map[string]any{"Id": uint64(1), "Name": "foo", "Age": uint(18)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.s, tt.args.removeZero, tt.args.deep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMapByJsonTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumberFieldMap(t *testing.T) {
	type args struct {
		list      any
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want map[uint]bool
	}{
		{
			name: "",
			args: args{
				list:      testUserList,
				fieldName: "Age",
			},
			want: map[uint]bool{
				18: true,
				25: true,
				30: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNumberFieldMap[uint](tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNumberMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringFieldMap(t *testing.T) {
	type args struct {
		list      any
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "",
			args: args{
				list:      testUserList,
				fieldName: "Name",
			},
			want: map[string]bool{
				testUserList[0].Name: true,
				testUserList[1].Name: true,
				testUserList[2].Name: true,
				testUserList[3].Name: true,
				testUserList[4].Name: true,
				testUserList[5].Name: true,
				testUserList[6].Name: true,
				testUserList[7].Name: true,
				testUserList[8].Name: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringFieldMap(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyBy(t *testing.T) {
	type args struct {
		list      any
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "",
			args: args{
				list:      testUserList,
				fieldName: "Id",
			},
			want: map[uint64]*testUser{
				1: testUserList[0],
				2: testUserList[1],
				3: testUserList[2],
				4: testUserList[3],
				5: testUserList[4],
				6: testUserList[5],
				7: testUserList[6],
				8: testUserList[7],
				9: testUserList[8],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyBy(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
