package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/handlers"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"sync"
	"time"
)

// client holds the limiter and lastSeen time for each user.
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimit restricts clients to 2 requests per second on average,
// with a maximum burst of 4 requests.
// It works per IP address.
// A 429 error is returned if the limit is exceeded.
func RateLimit() gin.HandlerFunc {
	var (
		mu      sync.Mutex
		clients = map[string]*client{}
	)

	// launch a goroutine to run the client cleanup every minute
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			cleanup(clients)
			mu.Unlock()
		}
	}()

	return func(ctx *gin.Context) {
		// extract client's IP address
		ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
		if err != nil {
			handlers.HandleInternalServerError(ctx, err)
			return
		}

		// add client to map if it doesn't exist
		// and update last seen
		mu.Lock()
		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastSeen = time.Now()
		mu.Unlock()

		// return a 429 error if the limit has been exceeded
		if !clients[ip].limiter.Allow() {
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

// cleans up clients older than 3 minutes
func cleanup(clients map[string]*client) {
	for ip, client := range clients {
		if time.Since(client.lastSeen) > 3*time.Minute {
			delete(clients, ip)
		}
	}
}
