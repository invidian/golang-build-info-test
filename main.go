package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("build info not available")
	}

	fmt.Println(bi.Main.Version)
}
