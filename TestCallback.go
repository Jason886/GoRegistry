package main

//import "fmt"

import "unsafe"

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: clib.a
#include "clib.h"
#include "clib_wrapper.h"
*/
import "C"

type hello_cb func(userdata unsafe.Pointer, a int) int;

type data_t struct {
	ud unsafe.Pointer
	cb hello_cb
}

//export hello_cb_go
func hello_cb_go(cb_wrap unsafe.Pointer, a int) int {
var cb = (*data_t)(cb_wrap)
var ret = cb.cb(cb.ud, a)
return ret
}

func hello_go(userdata unsafe.Pointer, a int, b int, cb hello_cb) {
var cb_wrap = &data_t{userdata, cb}
var aa C.int = C.int(a)
var bb C.int = C.int(b)
C.hello_wrapper(unsafe.Pointer(cb_wrap), aa, bb, C.hello_cb_wrapper)	
}

func my_callback(ud unsafe.Pointer, a int) int {
	println(a)
	return 0
}

func main() {
	var data = &struct {x int} {100}
	hello_go(unsafe.Pointer(data), 3, 5, my_callback)
}
