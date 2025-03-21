package pluck

import (
	"fmt"
	"reflect"

	"github.com/ml444/gutil/typex"
)

type TNumber interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

func ToNumbers[T TNumber](list any, fieldName string, unique bool) []T {
	var result []T
	vo := reflect.Indirect(reflect.ValueOf(list))
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				panic("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				panic(fmt.Sprintf("struct missed field %s", fieldName))
			}

			appendTo(f, &result)

		}
	default:
		panic("required list of struct type")
	}
	if unique {
		result = ToUniqueNumbers(result)
	}
	return result
}

func appendTo[T TNumber](f reflect.Value, result *[]T) {
	if f.Kind() == reflect.Ptr {
		if f.IsNil() {
			return
		}
		f = f.Elem()
	}
	switch f.Kind() {
	case reflect.Array, reflect.Slice:
		for j := 0; j < f.Len(); j++ {
			subF := f.Index(j)
			appendTo(subF, result)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		*result = append(*result, T(f.Uint()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		*result = append(*result, T(f.Int()))
	case reflect.Float32, reflect.Float64:
		*result = append(*result, T(f.Float()))
	default:
		return
	}
}

func ToStrings(list any, fieldName string) []string {
	var result []string
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				panic("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				panic(fmt.Sprintf("struct missed field %s", fieldName))
			}

			if f.Kind() != reflect.String {
				panic(fmt.Sprintf("struct element %s type required string", fieldName))
			}
			result = append(result, f.String())
		}
	default:
		panic("required list of struct type")
	}
	return result
}

func StructToValues(data any, omitempty bool) (values []any) {
	dataV := reflect.ValueOf(data)
	if dataV.Type().Kind() == reflect.Ptr {
		dataV = dataV.Elem()
	}

	for i := 0; i < dataV.NumField(); i++ {
		v := dataV.Field(i)
		if !v.IsValid() {
			continue
		}

		if v.IsZero() {
			if omitempty {
				continue
			}
			values = append(values, reflect.Zero(v.Type()).Interface())
		} else {
			values = append(values, v.Interface())
		}
	}

	return values
}

func StructToValueStrings(data any, omitempty bool) (values []string) {
	dataV := reflect.ValueOf(data)
	if dataV.Type().Kind() == reflect.Ptr {
		dataV = dataV.Elem()
	}

	for i := 0; i < dataV.NumField(); i++ {
		v := dataV.Field(i)

		// Check if the field is zero or empty
		isEmpty := reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
		if isEmpty && omitempty {
			continue
		}
		if isEmpty {
			values = append(values, typex.AnyToStr(reflect.Zero(v.Type()).Interface()))
		} else {
			values = append(values, typex.AnyToStr(v.Interface()))
		}
	}

	return values
}
