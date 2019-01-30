package common

import (
	"reflect"
	"unsafe"
)

const (
	SliceHeaderSize = int(unsafe.Sizeof(reflect.SliceHeader{}))
	StringHeaderSize = int(unsafe.Sizeof(reflect.StringHeader{}))

	SliceHeaderSizeUintPtr = unsafe.Sizeof(reflect.SliceHeader{})
	StringHeaderSizeUintPtr = unsafe.Sizeof(reflect.StringHeader{})
)
