package main

import (
	"fmt"
	"os"
)

func run(args []string) int {
	app := NewApp()

	if 1 == len(args) {
		args = append(args, "--help")
	}

	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args))
}
