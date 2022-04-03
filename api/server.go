package api

import (
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Serve runs Server on given port for the api.yml OpenAPI specification.
func Serve(port uint) {
	r := gin.Default()
	r.GET("/v1/diff", GetDiff)
	r.Run(fmt.Sprintf(":%d", port))
}

// GetDiff handles a GET request that computes a AbsoluteDateDifference on two query strings: start, end
func GetDiff(c *gin.Context) {
	start, end := c.Query("start"), c.Query("end")
	difference, err := diff.AbsoluteDateDifference(start, end)
	if err != nil {
		setErrorResponse(c, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": difference,
	})
}

func setErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Errorf("invalid request body. Error: %w", err).Error(),
	})
}
