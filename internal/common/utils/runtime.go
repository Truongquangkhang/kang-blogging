package utils

import "runtime"

// Retrieves the call stack, then retrieves the
// function name from function at depth-level
func GetFuncNameFromCallStack(depth int) string {
	pc := make([]uintptr, depth+1) // at least 1 entry need
	runtime.Callers(depth, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
