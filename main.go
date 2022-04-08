package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Args: %+v\n", os.Args)
	fmt.Printf("Env: %+v\n", os.Environ())
}
