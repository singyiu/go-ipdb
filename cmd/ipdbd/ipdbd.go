package main

import (
	ma "github.com/multiformats/go-multiaddr"
	"github.com/namsral/flag"
	"log"
	"os"
)

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "IPDB", 0)
	apiAddrStr := fs.String("apiAddr", "/ip4/127.0.0.1/tcp/6007", "gRPC API bind address")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	apiAddr, err := ma.NewMultiaddr(*apiAddrStr)
	if err != nil {
		log.Fatal(err)
	}

}
