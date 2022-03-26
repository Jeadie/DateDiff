package main

import (
	"fmt"
	"os"
)

type Subcommand string

const (
	CompareCmd Subcommand = "cmp"
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

	default:
		usage()
		return
	}
}

func usage() {
	fmt.Println("Hello world")
}
