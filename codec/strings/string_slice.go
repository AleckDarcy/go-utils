package strings

import (
	"github.com/AleckDarcy/go-utils/codec/common"
	"reflect"
	"unsafe"
)

func calStringSliceSize(slice []string) int {
	size := common.SliceHeaderSize + len(slice)*common.StringHeaderSize

	for _, str := range slice {
		size += len(str)
	}

	return size
}

func DeepCopyStringSlice(slice []string) []string {
	if slice == nil {
		return nil
	} else if len(slice) == 0 {
		return []string{}
	}

	bytes := make([]byte, calStringSliceSize(slice))
	globalStart := (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(globalStart))
	sliceHeader.Data = globalStart + common.SliceHeaderSizeUintPtr
	sliceHeader.Len = len(slice)
	sliceHeader.Cap = len(slice)

	stringHeaderArrayStart := sliceHeader.Data
	stringContentStart := stringHeaderArrayStart + uintptr(len(slice))*common.StringHeaderSizeUintPtr

	for _, str := range slice {
		stringHeader := (*reflect.StringHeader)(unsafe.Pointer(stringHeaderArrayStart))
		stringHeader.Data = stringContentStart
		stringHeader.Len = copy(bytes[stringContentStart-globalStart:], str)

		stringHeaderArrayStart += common.StringHeaderSizeUintPtr
		stringContentStart += uintptr(stringHeader.Len)
	}

	return *(*[]string)(unsafe.Pointer(globalStart))
}
