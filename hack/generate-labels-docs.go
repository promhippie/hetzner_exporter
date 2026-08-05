//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("docs/partials/labels.md")

	if err != nil {
		fmt.Printf("failed to create file")
		os.Exit(1)
	}

	defer f.Close()
}
