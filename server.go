package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Server(args []string) {
	router := gin.Default()
	fmt.Printf("%s\n", router.Routes())
}
