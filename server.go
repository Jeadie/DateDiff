package main

import (
	"flag"
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"github.com/gin-gonic/gin"
	"os"
)

func getDefaultPort() string {
	value, isSet := os.LookupEnv("PORT")
	if !isSet || len(value) == 0 {
		return "8000"
	}
	return value
}

func Server(args []string) {

	argFs := flag.NewFlagSet(string(ServerCmd), flag.ExitOnError)
	port := argFs.String("p", getDefaultPort(), "server connection port")
	argFs.Parse(args)

	_, err := diff.UintParse(*port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid port specified: %s\n", *port)
		return
	}

	gin.Default().Run(fmt.Sprintf(":%s", *port))
}
