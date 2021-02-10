package unsafebytes

import "unsafe"

// String converts a byte slice to a string in an unsafe, non-copy fashion. It
// works similarly to strings.Builder's String method.
func String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
