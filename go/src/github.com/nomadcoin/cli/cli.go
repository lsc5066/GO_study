package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
)

func usage() {

	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port1: Set the PORT of the rest API server\n\n")
	fmt.Printf("-port2: Set the PORT of the html server\n\n")
	fmt.Printf("-mode: Choose between 'html', 'rest' and 'both')\n\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set the PORT of the rest API server")
	mode := flag.String("mode", "rest", "Choose between 'html', 'rest' and 'both'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)

	case "html":
		explorer.Start(*port)

	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)

	default:
		usage()
	}
}
