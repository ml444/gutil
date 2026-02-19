package pluck

import (
	"reflect"
	"strings"
)

// DiffStruct returns fields in the target structure whose values ​​differ from those in the source data structure.
func DiffStruct(source, target interface{}, excludeFields ...string) map[string]interface{} {
	data := make(map[string]interface{})
	srcV := reflect.ValueOf(source)
	if srcV.Kind() == reflect.Ptr {
		srcV = srcV.Elem()
	}
	tgtV := reflect.ValueOf(target)
	if tgtV.Kind() == reflect.Ptr {
		tgtV = tgtV.Elem()
	}
	tgtT := tgtV.Type()
	for i := 0; i < tgtV.NumField(); i++ {
		v := tgtV.Field(i)
		t := tgtT.Field(i)
		if !t.IsExported() {
			continue
		}
		if v.Kind() == reflect.Ptr && v.IsNil() {
			continue
		}
		name := t.Name
		if Contains(excludeFields, name) {
			continue
		}
		jsonTag := t.Tag.Get("json")
		jsonTag = strings.Split(jsonTag, ",")[0]
		if jsonTag == "-" {
			continue
		}
		if jsonTag == "" {
			jsonTag = name
		}
		vv := srcV.FieldByName(name)
		if !vv.IsValid() {
			continue
		}
		if !reflect.DeepEqual(v.Interface(), vv.Interface()) {
			data[jsonTag] = v.Interface()
		}
	}
	return data
}

func Contains(source []string, name string) bool {
	for _, field := range source {
		if name == field {
			return true
		}
	}
	return false
}

type DigitAndString interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | string
}

// DiffArray returns the elements in target that are not in source.
func DiffArray[T DigitAndString](source, target []T) []T {
	var data []T
	m := make(map[T]struct{})
	for _, v := range source {
		m[v] = struct{}{}
	}
	for _, v := range target {
		if _, ok := m[v]; !ok {
			data = append(data, v)
		}
	}
	return data
}

// CompareList 比较两个切片或数组的元素是否一致，不考虑元素的顺序
func CompareList[T DigitAndString](source, target []T) bool {
	if len(source) != len(target) {
		return false
	}
	m := make(map[T]int)
	for _, v := range source {
		m[v]++
	}
	for _, v := range target {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
