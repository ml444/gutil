package debug

import "unsafe"

type iface struct {
	itab, data uintptr
}

func (i iface) Type() uintptr {
	return i.itab
}
func (i iface) Data() uintptr {
	return i.data
}

func GetTypePtr(i interface{}) uintptr {
	ins := *(*iface)(unsafe.Pointer(&i))
	return ins.Type()
}
func GetDataPtr(i interface{}) uintptr {
	ins := *(*iface)(unsafe.Pointer(&i))
	return ins.Data()
}
func GetTypeAndDataPtr(i interface{}) (uintptr, uintptr) {
	ins := *(*iface)(unsafe.Pointer(&i))
	return ins.Type(), ins.Data()
}
