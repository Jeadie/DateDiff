package main

import (
	"fmt"
	"os"
)

type Subcommand string

const (
	CompareCmd Subcommand = "cmp"
	DatasetCmd Subcommand = "dataset"
	ServerCmd  Subcommand = "server"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch Subcommand(os.Args[1]) {
	case CompareCmd:
		Cmp(os.Args[2:])
		return

	case ServerCmd:
		Server(os.Args[2:])
		return

	case DatasetCmd:
		Dataset(os.Args[2:])
		return

	default:
		usage()
		return
	}
}

func usage() {
	fmt.Println("Usage: datediff [SUBCOMMAND] <SUBCOMMAND ARGS>")
	fmt.Println("  Subcommands:")
	fmt.Println("    cmp: Compare two dates")
	fmt.Println("    server: Spin up a server that handles compare requests")
	fmt.Println("    dataset: Construct a dataset of input output pairs, with optionally, invalid inputs")
}
