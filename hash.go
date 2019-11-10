package hmap

import _ "unsafe"

//go:linkname _hash runtime.efaceHash
func _hash(v interface{}, seed uintptr) int

// hash wraps a general interface-hashing function
func hash(v interface{}) int {
	return _hash(v, 0)
}
