package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rmanzoku/apl"
)

func run() int {

	var apl apl.APL

	apl.CachePath = filepath.Join(os.Getenv("HOME"), ".cache", "apl")

	fmt.Println(apl)
	return 0
}

func main() {
	os.Exit(run())
}
