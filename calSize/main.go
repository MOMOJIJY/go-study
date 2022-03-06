package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	uid uint64
	username string
	tag int
}

type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
}

func mapSize() {
	key := uint64(1)
	value := User{}
	lenMap := int(99999)
	m := hmap{}
	size := int(unsafe.Sizeof(m))+lenMap*8*int(unsafe.Sizeof(key)+unsafe.Sizeof(value))
	fmt.Printf("map size = %v\n", size)
}

func main() {
	mapSize()
}