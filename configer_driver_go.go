package configer_driver_go

/*
#cgo LDFLAGS:  -lconfiger -lgdbm -lstdc++
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
	version = "1.0.0"
	codeSuccess = 0
)

var (
	_valLength = 3*1024*1024 // value size
)

// configer init
func ConfInit(path string,shmKey int,valLength int){
	// 安装路径
	configerPath := C.CString(path)
	C.configer_init(configerPath,C.int(shmKey))

	if valLength > 0 {
		_valLength = valLength
	}
}

// get value
func ConfGet(key string)(string,error){
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cPtrValue *C.char
	cValue := make([]C.char, _valLength)
	cPtrValue = (*C.char)(unsafe.Pointer(&(cValue[0])))

	// 共享内存获取
	shmRet := C.get_tbl_val(cKey,cPtrValue)
	if shmRet == codeSuccess { // 存在情况下
		return C.GoString(cPtrValue),nil
	}

	// 本地dump获取
	dumpRet := C.get_dump_tbl_val(cKey,cPtrValue)
	if dumpRet == codeSuccess {
		return C.GoString(cPtrValue),nil
	}

	errorMsg := fmt.Sprintf("configer go get %s is faild ,get shm code %d, get dump code %d",key,shmRet,dumpRet)
	return "",errors.New(errorMsg)
}
