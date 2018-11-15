package main

import (
	"github.com/the-code-innovator/go-block-chain/command"
	"os"
)

func main() {
	defer os.Exit(0)
	commandLineIterface := command.CommandLineInterface{}
	commandLineIterface.Run()
}
