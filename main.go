package main

import (
	"fmt"
	"os"
)

// Version is the version of this app. This value is injected by Makefile.
var Version = ""

func main() {
	for _, arg := range os.Args {
		switch arg {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		}
	}
}
