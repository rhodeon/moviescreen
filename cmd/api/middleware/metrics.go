package middleware

import (
	"expvar"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Metrics() gin.HandlerFunc {
	totalRequestsReceived := expvar.NewInt("total_requests_received")
	totalResponsesSent := expvar.NewInt("total_responses_sent")
	totalProcessingTimeInMicroseconds := expvar.NewInt("total_processing_time_microseconds")
	ResponseStatusCodeCount := expvar.NewMap("response_status_code_count")

	return func(ctx *gin.Context) {
		start := time.Now()

		// increment request count before proceeding
		totalRequestsReceived.Add(1)

		// proceed to next handler
		ctx.Next()

		// increment response count upon return from next handler
		totalResponsesSent.Add(1)

		// determine time taken for request handling
		duration := time.Now().Sub(start).Milliseconds()
		totalProcessingTimeInMicroseconds.Add(duration)

		// increment the count of the returned status code
		ResponseStatusCodeCount.Add(strconv.Itoa(ctx.Writer.Status()), 1)
	}
}
