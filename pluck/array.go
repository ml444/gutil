package pluck

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ml444/gutil/typex"
)

func ToUint64s(list interface{}, fieldName string) []uint64 {
	var result []uint64
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

			if f.Kind() != reflect.Uint64 {
				panic(fmt.Sprintf("struct element %s type required uint64", fieldName))
			}

			result = append(result, f.Uint())
		}
	default:
		panic("required list of struct type")
	}
	return result
}

func ToInt64s(list interface{}, fieldName string) []int64 {
	var result []int64
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

			if f.Kind() != reflect.Int64 {
				panic(fmt.Sprintf("struct element %s type required int64", fieldName))
			}

			result = append(result, f.Int())
		}
	default:
		panic("required list of struct type")
	}
	return result
}

func ToUint64Map(list interface{}, fieldName string) map[uint64]bool {
	out := ToUint64s(list, fieldName)
	res := map[uint64]bool{}
	for _, v := range out {
		res[v] = true
	}
	return res
}

func ToUint32s(list interface{}, fieldName string) []uint32 {
	var result []uint32
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

			if f.Kind() != reflect.Uint32 {
				panic(fmt.Sprintf("struct element %s type required uint32", fieldName))
			}
			result = append(result, uint32(f.Uint()))
		}
	default:
		panic("required list of struct type")
	}
	return result
}
func ToUint32Map(list interface{}, fieldName string) map[uint32]bool {
	out := ToUint32s(list, fieldName)
	res := map[uint32]bool{}
	for _, v := range out {
		res[v] = true
	}
	return res
}
func ToString(list interface{}, fieldName string) []string {
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
func ToStringMap(list interface{}, fieldName string) map[string]bool {
	out := ToString(list, fieldName)
	res := map[string]bool{}
	for _, v := range out {
		res[v] = true
	}
	return res
}
func SubSlice(obj interface{}, startIdx, endIdx interface{}) (subSlice interface{}) {
	vo := reflect.ValueOf(obj)
	vk := vo.Kind()
	if vk != reflect.Slice && vk != reflect.Array {
		panic("obj required slice or array type")
	}
	list := reflect.MakeSlice(reflect.SliceOf(vo.Type().Elem()), 0, 0)
	subSlice = list.Interface()
	start, end := typex.AnyToInt(startIdx), typex.AnyToInt(endIdx)
	if start < 0 || end < 0 || start >= end {
		return
	}
	length := vo.Len()
	for i := start; i < length && i < end; i++ {
		list = reflect.Append(list, vo.Index(i))
	}
	subSlice = list.Interface()
	return
}

// KeyBy0 list is []StructType, return: *map[fieldType]StructType
func KeyBy0(list interface{}, fieldName string) (res interface{}) {
	// 取下 field type
	vo := typex.EnsureIsSliceOrArray(list)
	elType := vo.Type().Elem()
	t := elType
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("slice or array element required struct type, but got %v", t))
	}
	var keyType reflect.Type
	if sf, ok := t.FieldByName(fieldName); ok {
		keyType = sf.Type
	} else {
		panic(fmt.Sprintf("not found field %s", fieldName))
	}
	m := reflect.MakeMap(reflect.MapOf(keyType, elType))
	//resVo := reflect.ValueOf(res)
	//if resVo.Kind() != reflect.Ptr {
	//	panic(fmt.Sprintf("invalid res type %v, required *map[key]val", resVo.Type()))
	//}
	//resVo = resVo.Elem()
	//typex.EnsureIsMapType(resVo, keyType, elType)
	l := vo.Len()
	for i := 0; i < l; i++ {
		el := vo.Index(i)
		elDef := el
		for elDef.Kind() == reflect.Ptr {
			elDef = elDef.Elem()
		}
		f := elDef.FieldByName(fieldName)
		if !f.IsValid() {
			continue
		}
		m.SetMapIndex(f, el)
	}
	//resVo.Set(m)
	return m.Interface()
}

// KeyBy 根据KeyFieldName 取出 map[key]val
// 允许list长度为0
// 相对于 KeyBy 更加安全，且可以适应 Key 为nil的场景
// 性能方便相较 KeyBy 版本有大幅度提升，list的长度越长，性能提升越明显，根据性能测试，CPU/内存消耗相对较少且内存分配次数也大幅度减少
func KeyBy(list interface{}, fieldName string) interface{} {
	lv := reflect.ValueOf(list)
	// if lv.IsNil() {
	// 	return reflect.New(lv.Type()).Elem().Interface()
	// }

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

// SliceToUint64Map list 转化 map ，用于判断某个值是否存在 list 中
func SliceToUint64Map(list []uint64) map[uint64]bool {
	result := map[uint64]bool{}
	for _, v := range list {
		result[v] = true
	}
	return result
}

func JoinIntegerSliceToString(list interface{}, sep string) (string, error) {
	listValueReflect := reflect.ValueOf(list)
	switch listValueReflect.Index(0).Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		var intList []int
		for i := 0; i < listValueReflect.Len(); i++ {
			intList = append(intList, typex.AnyToInt(listValueReflect.Index(i).Interface()))
		}

		var intSlice2String string
		for _, item := range intList {
			intSlice2String = intSlice2String + strconv.Itoa(item) + sep
		}
		return strings.TrimRight(intSlice2String, sep), nil
	default:
		return "", errors.New(fmt.Sprintf("Argument 1 expect integer type, given %t", list))
	}
}

func StructToSlice(data interface{}, omitempty bool) (columns []interface{}) {
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
			columns = append(columns, reflect.Zero(v.Type()).Interface())
		} else {
			columns = append(columns, v.Interface())
		}
	}

	return columns
}

func StructToStrSlice(data interface{}, omitempty bool) (columns []string) {
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
			columns = append(columns, typex.AnyToStr(reflect.Zero(v.Type()).Interface()))
		} else {
			columns = append(columns, typex.AnyToStr(v.Interface()))
		}
	}

	return columns
}
