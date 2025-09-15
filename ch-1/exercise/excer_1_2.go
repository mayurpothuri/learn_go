package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		fmt.Println("index:", k, "value:", v)
	}
}
