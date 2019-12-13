package configer_go

/*
#cgo LDFLAGS:  -lconfiger -L/usr/local/lib   -lgdbm -lstdc++
#include <stdlib.h>
#include <stdio.h>
#include "api.h"
*/
import "C"

import (
	"unsafe"
)

var (
	_valLength = 3*1024*1024 // value 大小
)

// 配置初始化
func ConfigerInit(path string,valLength int){
	// 安装路径
	configerPath := C.CString(path)
	C.configer_init(configerPath)

	if valLength > 0 {
		_valLength = valLength
	}
}

// etcd获取
func EtcdGet(key string)(string,error){
	etcdKey := C.CString(key);
	defer C.free(unsafe.Pointer(etcdKey))
	var cPtrValue *C.char
	cValue := make([]C.char, _valLength)
	cPtrValue = (*C.char)(unsafe.Pointer(&(cValue[0])))

	// 共享内存获取
	C.get_tbl_val(etcdKey,cPtrValue);
	etcdValue := C.GoString(cPtrValue)
	if len(etcdValue) <= 0 {
		// dump获取
		C.get_dump_tbl_val(etcdKey,cPtrValue);
		etcdValue = C.GoString(cPtrValue)
	}
	return etcdValue,nil
}
