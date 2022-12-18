package main

import (
	"github.com/neglect-yp/enumizer/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
