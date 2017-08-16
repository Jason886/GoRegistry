package main

import "./Registry"
import "unsafe"
import "fmt"

type Data struct {
    x int
}

func main() {
    var aa = &Data {100}
    var bb [1000]uint64
    for i:=0; i<1000; i++ {
	id, ok := Registry.Ref(unsafe.Pointer(aa))
        fmt.Printf("%X,", id)
	println(id, ok)
        bb[i] = id
    }
    for i:=900; i<1000; i++ {
	Registry.Unref(bb[i])    
}
    for i:=0; i<1000; i++ {
	var obj = (*Data)(Registry.Get(bb[i]))
        fmt.Printf("%#v\n", obj)
    }
    
}
