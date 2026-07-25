package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("devopshub-cli version %s\n", version)
		return
	}

	fmt.Println("DevOpsHub CLI - Open-source DevOps toolkit")
	fmt.Println("Usage: devopshub <command> [flags]")
	fmt.Println("\nRun 'devopshub --help' for more details.")
}
