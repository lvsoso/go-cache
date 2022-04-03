package cache

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I${SRCDIR}/../../rocksdb/include
// #cgo LDFLAGS: -L${SRCDIR}/../../rocksdb -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -ldl -O3
import "C"
import (
	"errors"
	"unsafe"
)

func (c *rocksdbCache) Set(key string, value []byte) error {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	v := C.CBytes(value)
	defer C.free(v)
	C.rocksdb_put(c.db, c.wo, k, C.size_t(len(key)), (*C.char)(v), C.size_t(len(value)), &c.e)
	if c.e != nil {
		return errors.New(C.GoString(c.e))
	}
	return nil
}
