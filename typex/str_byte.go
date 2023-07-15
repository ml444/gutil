package typex

import (
	"unsafe"
)

// Str2Byte converts string to a byte slice without memory allocation.
func Str2Byte(s string) []byte {
	/* #nosec G103 */
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Byte2Str converts byte slice to a string without memory allocation.
func Byte2Str(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}
