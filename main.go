package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/naoty/mdserve/server"
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

	logger := log.New(os.Stdout, "", 0)
	server := server.New().WithLogger(logger)
	http.ListenAndServe(":8000", server)
}
