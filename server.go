package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)
const DefaultPort = uint(8000)


func Server(args []string) {

	argFs := flag.NewFlagSet(string(ServerCmd), flag.ExitOnError)
	port := argFs.Uint("p", DefaultPort, "server connection port")
	argFs.Parse(args)

	gin.Default().Run(fmt.Sprintf(":%d", *port))
}
