package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"golang.org/x/time/rate"
	"net/http"
)

// RateLimit restricts clients to 2 requests per second on average,
// with a maximum burst of 4 requests.
// It returns a 429 error if the limit is exceeded.
func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				response.ErrorResponse(
					http.StatusTooManyRequests,
					response.GenericError(response.ErrRateLimitExceeded),
				),
			)
			return
		}
		ctx.Next()
	}
}
