package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Timestamp   string `json:"timestamp"`
	Hostname    string `json:"hostname"`
	Method      string `json:"method"`
}

func main() {
	router := gin.Default()
	router.GET("/appRoundRobin", getApp)
	router.GET("/leastConnection", getApp)
	router.GET("/ipHash", getApp)

	router.Run("localhost:8080")
}

func getApp(c *gin.Context) {
	t := time.Now().Format(time.Kitchen)
	// Note: without explicit zone, returns time in given location.

	var response = []Response{
		{Description: fmt.Sprintf("Hi, this is from app %s", os.Args), Name: "Blue Train", Timestamp: t, Hostname: c.Request.RemoteAddr, Method: c.Request.Method},
	}
	c.IndentedJSON(http.StatusOK, response)
}
