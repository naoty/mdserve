package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/naoty/mdserve/contents"
	"github.com/naoty/mdserve/server"
	"github.com/spf13/pflag"
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
	helpFlag := pflag.BoolP("help", "h", false, "")
	versionFlag := pflag.BoolP("version", "v", false, "")
	pflag.Parse()

	if *helpFlag {
		fmt.Println(help)
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	if pflag.NArg() == 0 {
		fmt.Println(help)
		os.Exit(1)
	}

	dir := pflag.Arg(0)
	info, err := os.Stat(dir)
	if err != nil {
		log.Fatalln(err)
	}
	if !info.IsDir() {
		fmt.Println(help)
		os.Exit(1)
	}

	err = filepath.Walk(dir, parse)
	if err != nil {
		log.Fatalln(err)
	}

	logger := log.New(os.Stdout, "", 0)
	server := server.New(dir).WithLogger(logger)
	http.ListenAndServe(":8000", server)
}

func parse(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || filepath.Ext(path) != ".md" {
		return nil
	}

	return contents.Parse(path)
}
