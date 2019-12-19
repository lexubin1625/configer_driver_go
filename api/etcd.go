package driver

/*
#cgo LDFLAGS:  -lconfiger -L/usr/local/lib   -lgdbm -lstdc++ 
#include <stdlib.h>
#include <stdio.h>
#include "configer/api.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	success_code = 0
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
	etcdKey := C.CString(key)
	defer C.free(unsafe.Pointer(etcdKey))
	var cPtrValue *C.char
	cValue := make([]C.char, _valLength)
	cPtrValue = (*C.char)(unsafe.Pointer(&(cValue[0])))

	// 共享内存获取
	shmRet := C.get_tbl_val(etcdKey,cPtrValue)
	if shmRet == success_code { // 存在情况下
		return C.GoString(cPtrValue),nil
	}

	// 本地dump获取
	dumpRet := C.get_dump_tbl_val(etcdKey,cPtrValue)
	if dumpRet == success_code {
		return C.GoString(cPtrValue),nil
	}

	errorMsg := fmt.Sprintf("configer go get %s is faild ,get shm code %d, get dump code %d",key,shmRet,dumpRet)
	return "",errors.New(errorMsg)
}
