package main

import (
	"fmt"
	"os"

	"github.com/hamsa0x7/strand/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
