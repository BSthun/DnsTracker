package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Log(fn logger, message string) {
	_, file, line, ok := runtime.Caller(1)
	dir := strings.Split(filepath.Dir(file), "/")
	if ok {
		fn("[" + dir[len(dir)-1] + "/" + filepath.Base(file) + ":" + strconv.Itoa(line) + "] " + message)
	} else {
		fn("[N/A] " + message)
	}
}
