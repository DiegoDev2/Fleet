package main

import (
	"github.com/DiegoDev2/Fleet/internal/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
