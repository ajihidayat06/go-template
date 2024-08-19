package common

import (
	"fmt"
	"runtime"
	"strings"
)

func GetRuntimeCaller(skip int) string {
	pc, fileName, line, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc)

	// Mengambil nama fungsi dan memisahkan dari path package dan receiver
	fullFnName := fn.Name()
	parts := strings.Split(fullFnName, ".")
	fnName := parts[len(parts)-1] // Ambil bagian terakhir, yang biasanya adalah nama fungsinya

	return fmt.Sprintf("%s>%s:%d", fileName, fnName, line)
}
