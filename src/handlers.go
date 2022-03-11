package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RelayHandler(c *gin.Context) {
	relay := c.Param("id")
	state := c.Query("state")
	c.IndentedJSON(http.StatusOK, gin.H{"relay": relay, "state": state})
}
