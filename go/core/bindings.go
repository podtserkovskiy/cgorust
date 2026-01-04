package core

// #include "rust/core.h"
import "C"

// Add wraps the Rust add function
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
