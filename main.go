package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lenz-gohash/tcp/client"
	"github.com/lenz-gohash/tcp/server"
)

// Parse the command line arguments

func main() {
	const serverAddress = ":12345"

	argc := len(os.Args)
	argv := os.Args

	// Usage message
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [server | client] \n", argv[0])
		flag.PrintDefaults()
	}

	if argc < 2 {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	if argv[1] == "server" {
		server.Serve(serverAddress)
	} else if argv[1] == "client" {
		client.Dial(serverAddress, "Hello! World!")
	} else {
		flag.CommandLine.Usage()
		os.Exit(1)
	}
}
