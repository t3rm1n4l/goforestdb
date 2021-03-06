package forestdb

//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

//#include <stdlib.h>
//#include <libforestdb/forestdb.h>
import "C"

import (
	"unsafe"
)

// GetKV simplified API for key/value access to Get()
func (k *KVStore) GetKV(key []byte) ([]byte, error) {

	var kk unsafe.Pointer
	if len(key) != 0 {
		kk = unsafe.Pointer(&key[0])
	}
	lenk := len(key)

	var bodyLen C.size_t
	var bodyPointer unsafe.Pointer

	errNo := C.fdb_get_kv(k.db, kk, C.size_t(lenk), &bodyPointer, &bodyLen)
	if errNo != RESULT_SUCCESS {
		return nil, Error(errNo)
	}

	body := C.GoBytes(bodyPointer, C.int(bodyLen))
	C.free(bodyPointer)
	return body, nil
}

// SetKV simplified API for key/value access to Set()
func (k *KVStore) SetKV(key, value []byte) error {

	var kk, v unsafe.Pointer

	if len(key) != 0 {
		kk = unsafe.Pointer(&key[0])
	}

	if len(value) != 0 {
		v = unsafe.Pointer(&value[0])
	}

	lenk := len(key)
	lenv := len(value)

	errNo := C.fdb_set_kv(k.db, kk, C.size_t(lenk), v, C.size_t(lenv))
	if errNo != RESULT_SUCCESS {
		return Error(errNo)
	}
	return nil
}

// DeleteKV simplified API for key/value access to Delete()
func (k *KVStore) DeleteKV(key []byte) error {

	var kk unsafe.Pointer
	if len(key) != 0 {
		kk = unsafe.Pointer(&key[0])
	}

	lenk := len(key)

	errNo := C.fdb_del_kv(k.db, kk, C.size_t(lenk))
	if errNo != RESULT_SUCCESS {
		return Error(errNo)
	}
	return nil
}
