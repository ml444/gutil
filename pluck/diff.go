package pluck

import (
	"reflect"
	"strings"
)

func DiffStruct(source, target interface{}, excludeFields []string) map[string]interface{} {
	var data = make(map[string]interface{})
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
			data[jsonTag] = vv.Interface()
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
