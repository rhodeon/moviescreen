package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/cmd/api/responseErrors"
	"github.com/rhodeon/prettylog"
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

// RateLimit restricts clients to the configured number of
// requests-per-second and maximum burst.
// It works per IP address.
// A 429 error is returned if the limit is exceeded.
func RateLimit(config common.Config) gin.HandlerFunc {
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
		// carry out check if rate limiting is enabled
		if config.Limiter.Enabled {
			// extract client's IP address
			ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
			if err != nil {
				prettylog.ErrorF("internal server error: %s", err.Error())
				responseErrors.HandleInternalServerError(ctx, err)
				return
			}

			// add client to map if it doesn't exist
			// and update last seen
			mu.Lock()
			if _, exists := clients[ip]; !exists {
				clients[ip] = &client{
					limiter: rate.NewLimiter(
						rate.Limit(config.Limiter.Rps),
						config.Limiter.Burst,
					),
				}
			}
			clients[ip].lastSeen = time.Now()
			mu.Unlock()

			// return a 429 error if the limit has been exceeded
			if !clients[ip].limiter.Allow() {
				responseErrors.SetStatusAndBody(
					ctx,
					http.StatusTooManyRequests,
					response.GenericError(responseErrors.ErrMessageRateLimitExceeded),
				)
				return
			}
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
