package middleware

import (
	"github.com/gin-gonic/gin"
)

// HandleErrors handle error to prepare error response
func HandleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	lastError := c.Errors.Last()

	c.IndentedJSON(-1, lastError)
}
