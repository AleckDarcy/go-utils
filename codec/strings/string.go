package strings

import (
	"reflect"
	"unsafe"
)

func DeepCopyString(str string) string {
	bytes := []byte(str)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

	stringHeader := &reflect.StringHeader{
		Data: sliceHeader.Data,
		Len: sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(stringHeader))
}
