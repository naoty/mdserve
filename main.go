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
  mdserve <dir> [-H | --host <hostname>] [-p | --port <port>]
  mdserve -h | --help
  mdserve -v | --version

Options:
  -H --host <hostname> Run web server at specified hostname [default: localhost].
  -p --port <port>     Run web server at specified port [default: 8000].
  -h --help            Show this message.
  -v --version         Show version.`

func main() {
	host := pflag.StringP("host", "H", "localhost", "")
	port := pflag.IntP("port", "p", 8000, "")
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

	addr := fmt.Sprintf("%s:%d", *host, *port)
	logger.Printf("listening on %s\n", addr)

	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatalln(err)
	}
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
