package main

import "fmt"
import _ "embed"
import "runtime/debug"


//go:embed VERSION
var version string

func main() {
	fmt.Println(version)

	info, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Println(info.Main.Version)
	}
}
