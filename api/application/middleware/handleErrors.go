package middleware

import "github.com/gin-gonic/gin"

// HandleErrors handle error to prepare error response
func HandleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		//c.JSON(-1, c.Errors) // -1 == not override the current error code
	}
}
