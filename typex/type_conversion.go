package typex

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

func AnyToByte(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func AnyToInt(i interface{}) int {
	switch x := i.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case string:
		i64, _ := strconv.ParseInt(x, 10, 64)
		return int(i64)
	}
	return 0
}

func AnyToStr(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case *string:
		return *x
	case int:
		return strconv.Itoa(x)
	case int8:
		return strconv.Itoa(int(x))
	case int16:
		return strconv.Itoa(int(x))
	case int32:
		return strconv.Itoa(int(x))
	case int64:
		return strconv.FormatInt(x, 10)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case uint64:
		return strconv.FormatUint(x, 10)
	case float32:
		return strconv.FormatFloat(float64(x), 'f', 6, 32)
	case float64:
		return strconv.FormatFloat(x, 'f', 6, 64)
	case bool:
		return strconv.FormatBool(x)
	case []byte:
		return string(x)
	case *[]byte:
		return string(*x)
	default:
		buf, err := json.Marshal(x)
		if err != nil {
			var b strings.Builder
			encoder := gob.NewEncoder(&b)
			err = encoder.Encode(x)
			if err == nil {
				return b.String()
			}
		} else {
			return string(buf)
		}
	}
	return ""
}
