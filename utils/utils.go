package utils

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadFile(pathFromCaller string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of uitls.ReadFile")
	}
	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	f, err := os.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}
	str := string(f[:])
	str = strings.TrimRight(str, "\n")
	return str
	// return strings.Split(str, "\n")
}
