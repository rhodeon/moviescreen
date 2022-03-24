package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MaxSizeLimit restricts the size of requests to 1MB.
func MaxSizeLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		maxBytes := 1_048_576 // max size of 1MB

		// any future attempts to read a body greater than 1MB will return an error
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(maxBytes))

		ctx.Next()
	}
}
