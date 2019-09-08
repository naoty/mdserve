package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/naoty/mdserve/contents"
	"github.com/naoty/mdserve/server"
)

// Version is the version of this app. This value is injected by Makefile.
var Version = ""

var help = `Usage:
  mdserve <dir>
  mdserve -h | --help
  mdserve -v | --version

Options:
  -h --help     Show this message.
  -v --version  Show version.`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help)
		os.Exit(1)
	}

	dir := ""
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			fmt.Println(help)
			os.Exit(0)
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		default:
			info, err := os.Stat(arg)
			if err != nil {
				log.Fatalln(err)
			}

			if !info.IsDir() {
				fmt.Println(help)
				os.Exit(1)
			}

			dir = arg
		}
	}

	err := filepath.Walk(dir, parse)
	if err != nil {
		log.Fatalln(err)
	}

	logger := log.New(os.Stdout, "", 0)
	server := server.New().WithLogger(logger)
	http.ListenAndServe(":8000", server)
}

func parse(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || filepath.Ext(path) != ".md" {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = contents.Parse(file)
	if err != nil {
		file.Close()
		return err
	}

	return nil
}
