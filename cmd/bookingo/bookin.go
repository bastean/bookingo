package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bastean/bookingo/pkg/cmd/server"
)

const cli = "bookingo"

var port string

func usage() {
	fmt.Printf("Usage: %v [OPTIONS]\n", cli)
	fmt.Printf("\nE.g.: %v -p 8080\n\n", cli)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&port, "p", os.Getenv("PORT"), "Port")

	flag.Usage = usage

	flag.Parse()

	server.Run(port)
}
