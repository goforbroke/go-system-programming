package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "%d\n", 10)
	fmt.Fprintf(f, "%.2f\n", 10.1)
	fmt.Fprintf(f, "%s\n", "string")
}
