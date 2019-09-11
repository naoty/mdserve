package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/naoty/mdserve/contents"
	"github.com/naoty/mdserve/server"
	"github.com/spf13/pflag"
)

// Version is the version of this app. This value is injected by Makefile.
var Version = ""

var help = `Usage:
  mdserve <dir> [-H | --host <hostname>] [-p | --port <port>] [-w | --watch]
  mdserve -h | --help
  mdserve -v | --version

Options:
  -H --host <hostname>  Run web server at specified hostname [default: localhost].
  -p --port <port>      Run web server at specified port [default: 8000].
  -h --help             Show this message.
  -v --version          Show version.
  -w --watch            Watch changes in contents and parse them.`

func main() {
	host := pflag.StringP("host", "H", "localhost", "")
	port := pflag.IntP("port", "p", 8000, "")
	helpFlag := pflag.BoolP("help", "h", false, "")
	versionFlag := pflag.BoolP("version", "v", false, "")
	watchFlag := pflag.BoolP("watch", "w", false, "")
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

	err = contents.ParseAll(dir)
	if err != nil {
		log.Fatalln(err)
	}

	if *watchFlag {
		watcherLogger := log.New(os.Stdout, "[watcher] ", 0)
		watcher, err := contents.NewWatcher(watcherLogger)
		if err != nil {
			log.Fatalln(err)
		}
		defer watcher.Close()

		err = watcher.AddAll(dir)
		if err != nil {
			watcher.Close()
			log.Fatalln(err)
		}

		watcher.Start()
	}

	logger := log.New(os.Stdout, "[server] ", 0)
	server := server.New(dir).WithLogger(logger)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	logger.Printf("listening on %s\n", addr)

	err = http.ListenAndServe(addr, server)
	if err != nil {
		log.Fatalln(err)
	}
}
