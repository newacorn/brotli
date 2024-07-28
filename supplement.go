package brotli

import "unsafe"

//go:noescape
//go:linkname MemClrNoHeapPointers runtime.memclrNoHeapPointers
//goland:noinspection GoUnusedParameter
func MemClrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
