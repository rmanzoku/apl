package main

import (
	"fmt"
	"os"
)

func run() int {

	fmt.Println("Hello apl")
	return 0
}

func main() {
	os.Exit(run())
}
