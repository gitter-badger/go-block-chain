package main

import (
	"os"

	"github.com/the-code-innovator/go-block-chain/command"
)

func main() {
	defer os.Exit(0)
	commandLineIterface := command.CommandLineInterface{}
	commandLineIterface.Run()
}
