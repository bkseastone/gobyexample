package main

import (
	"reflect"
	"unsafe"
)

func str2Bytes(str string) []byte {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	header.Len = len(str)
	header.Cap = header.Len
	return *(*[]byte)(unsafe.Pointer(header))
}
func bytes2Str(bts []byte) string {
	return *(*string)(unsafe.Pointer(&bts))
}
