package api

import (
	"fmt"
	"github.com/Jeadie/DateDiff/diff"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve(port uint) {
	r := gin.Default()
	r.GET("/v1/diff", GetDiff)
	r.Run(fmt.Sprintf(":%d", port))
}

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
