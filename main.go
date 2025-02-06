package main

import "fmt"
import _ "embed"

//go:embed VERSION
var version string

func main() {
	fmt.Println(version)
}
