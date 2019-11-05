package ray

/*
#cgo CFLAGS: -std=gnu99 -Wno-missing-braces -Wno-unused-result -Wno-unused-function -Wno-unused-variable -Wno-unused-local-typedef -Wno-macro-redefined
*/
import "C"

import "runtime"

func init() {
	// Make sure the main goroutine is bound to the main thread.
	runtime.LockOSThread()
}