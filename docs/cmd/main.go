package main

import (
	_ "embed"
	"os"
	"slices"
)

//go:embed htmx.min.js
var htmxJS []byte

func main() {
	if slices.Contains(os.Args, "static") {
		generate()
	} else {
		serve()
	}
}
