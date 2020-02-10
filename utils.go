package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetApplicationPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		return filepath.Dir(dir)
	}
	return ""
}

func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024, m.HeapReleased/1024)
}
