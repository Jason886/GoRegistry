package Registry

import "unsafe"
import "time"

var _last_ms uint64 = 0
var _index uint32 = 0
var _index_max uint32 = 0xFFFFF
var _map map[uint64] unsafe.Pointer = make(map[uint64] unsafe.Pointer, 512)

// TODO:ADD LOCK

func _get_uuid() (uint64, bool) {
    var err bool = false
    var ms uint64 = uint64( time.Now().UnixNano()/1e6)
    ms <<= 20
    ms &^= 0xFFFFF
    
    if _last_ms != ms {
        _last_ms = ms
        _index = 0
    }

    var uuid uint64 = ms + uint64(_index)
    for _, ok := _map[uuid]; ok; {
        _index++
        if _index > _index_max {
            _index = 0
            err = true
            break
        }
        uuid = ms + uint64(_index)
    }

    if !err {
        _index++
        return uuid, !err
    } else {
        return 0, !err
    }
}

func Ref(obj unsafe.Pointer) (uint64, bool) {
    id, ok := _get_uuid()
    if ok {
        _map[id] = obj
    }
    return id, ok
}

func Get(id uint64) unsafe.Pointer {
    obj, ok := _map[id]
    if ok {
        return obj
    }
    return nil
}

func Unref(id uint64) {
    delete(_map, id)
}
