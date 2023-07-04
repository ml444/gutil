package typex

/*


//go:build go1.20
// +build go1.20
*/

/*
copy from github.com/valyala/fasthttp
*/

//package typex
//
//import "unsafe"
//
//// Str2Byte converts string to a byte slice without memory allocation.
//func Str2Byte(s string) []byte {
//	return unsafe.Slice(unsafe.StringData(s), len(s))
//}
//
//// Byte2Str converts byte slice to a string without memory allocation.
//// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//func Byte2Str(b []byte) string {
//	return unsafe.String(unsafe.SliceData(b), len(b))
//}
