package pluck

import (
	"fmt"
	"reflect"
	"strings"
)

// 把结构体的字段构成map输出，如果存在json标签，则key使用json标签名，否则用字段自身的命名，value为字段值，如果字段值为nil，则不输出。
// removeZero：是否移除零值
// deep：是否深度转换，字段为结构体时，可以继续转换为map
func StructToMapByJsonTag(s any, removeZero bool, deep bool) map[string]any {
	m := make(map[string]any)
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return m
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		//if !v.Type().Field(i).IsExported() {
		//	continue
		//}
		if !field.IsValid() || !field.CanInterface() {
			continue
		}
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		if removeZero && field.IsZero() {
			continue
		}
		// json tag
		jsonTag := v.Type().Field(i).Tag.Get("json")
		var key string
		if jsonTag != "" && jsonTag != "-" {
			key = strings.Split(jsonTag, ",")[0]
		} else {
			key = v.Type().Field(i).Name
		}
		if deep && field.Kind() == reflect.Struct {
			m[key] = StructToMapByJsonTag(field.Interface(), removeZero, deep)
		} else {
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
			m[key] = field.Interface()
		}
	}
	return m
}

func ToNumberMap[T TNumber](list any, fieldName string) map[T]bool {
	out := ToNumbers[T](list, fieldName, false)
	res := map[T]bool{}
	for _, v := range out {
		res[v] = true
	}
	return res
}

func ToStringMap(list any, fieldName string) map[string]bool {
	out := ToStrings(list, fieldName)
	res := map[string]bool{}
	for _, v := range out {
		res[v] = true
	}
	return res
}

// KeyBy list is []StructType, return: map[fieldType]StructType
// KeyBy 根据KeyFieldName 取出 map[fieldType]StructType
// 允许list长度为0
// 相对于 KeyBy0 更加安全，且可以适应 Key 为nil的场景
// 性能方便相较 KeyBy0 版本有大幅度提升，list的长度越长，性能提升越明显，根据性能测试，CPU/内存消耗相对较少且内存分配次数也大幅度减少
func KeyBy(list any, fieldName string) any {
	lv := reflect.Indirect(reflect.ValueOf(list))

	switch lv.Type().Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		panic("element not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m.Interface()
}
