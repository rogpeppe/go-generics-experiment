// Code generated by go2go; DO NOT EDIT.


//line loadptr_test.go2:1
package gatomic

//line loadptr_test.go2:1
import (
//line loadptr_test.go2:1
 "sync/atomic"
//line loadptr_test.go2:1
 "testing"
//line loadptr_test.go2:1
 "unsafe"
//line loadptr_test.go2:1
)

//line loadptr_test.go2:7
func TestLoadStorePointer(t *testing.T) {
				var x int
				var p *int
//line loadptr_test.go2:9
  instantiate୦୦StorePointer୦int(&p, &x)
//line loadptr_test.go2:11
 pp := instantiate୦୦LoadPointer୦int(&p)
	*pp = 12
	if x != 12 {
		t.Fatal("unexpected value")
	}
}
//line loadptr.go2:12
func instantiate୦୦StorePointer୦int(addr **int, val *int,) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(addr)), unsafe.Pointer(val))
}
//line loadptr.go2:8
func instantiate୦୦LoadPointer୦int(addr **int,) *int {
	return (*int)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(addr))))
}

//line loadptr.go2:10
var _ = atomic.AddInt32
//line loadptr.go2:10
var _ = testing.AllocsPerRun

//line loadptr.go2:10
type _ unsafe.Pointer