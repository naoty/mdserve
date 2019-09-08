package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/naoty/mdserve/contents"
	"github.com/naoty/mdserve/server"
)

// Version is the version of this app. This value is injected by Makefile.
var Version = ""

var help = `Usage:
  mdserve -h | --help
  mdserve -v | --version

Options:
  -h --help     Show this message.
  -v --version  Show version.`

func main() {
	for _, arg := range os.Args {
		switch arg {
		case "-h", "--help":
			fmt.Println(help)
			os.Exit(0)
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		}
	}

	file, err := os.Open("examples/post.md")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	err = contents.Parse(file)
	if err != nil {
		file.Close()
		log.Fatalln(err)
	}

	logger := log.New(os.Stdout, "", 0)
	server := server.New().WithLogger(logger)
	http.ListenAndServe(":8000", server)
}
