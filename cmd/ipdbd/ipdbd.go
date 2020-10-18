package main

import (
	"github.com/namsral/flag"
	"log"
	"os"
)

func printHelp() {
	helpStr := `./ipdb -cmd help`
	log.Printf("%v", helpStr)
}

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "IPDB", 0)
	//apiAddrStr := fs.String("apiAddr", "/ip4/127.0.0.1/tcp/6007", "gRPC API bind address")
	cmdStr := fs.String("cmd", "help", "print help")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	/*
	apiAddr, err := ma.NewMultiaddr(*apiAddrStr)
	if err != nil {
		log.Fatal(err)
	}
	*/

	switch *cmdStr {
	case "help":
		printHelp()
	default:
		log.Printf("Err: unsupported cmd: %+v", cmdStr)
	}
}
