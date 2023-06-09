package typex

import (
	"fmt"
	"reflect"
)

func EnsureIsSliceOrArray(obj interface{}) (res reflect.Value) {
	vo := reflect.ValueOf(obj)
	for vo.Kind() == reflect.Ptr || vo.Kind() == reflect.Interface {
		vo = vo.Elem()
	}
	k := vo.Kind()
	if k != reflect.Slice && k != reflect.Array {
		panic(fmt.Sprintf("obj required slice or array type, but got %v", vo.Type()))
	}
	res = vo
	return
}

func EnsureIsMapType(m reflect.Value, keyType, valType reflect.Type) {
	if m.Kind() != reflect.Map {
		panic(fmt.Sprintf("required map type, but got %v", m.Type()))
	}
	t := m.Type()
	if t.Key() != keyType {
		panic(fmt.Sprintf("map key type not equal, %v != %v", t.Key(), keyType))
	}
	if t.Elem() != valType {
		panic(fmt.Sprintf("map val type not equal, %v != %v", t.Elem(), valType))
	}
}
