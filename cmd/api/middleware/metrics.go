package middleware

import (
	"expvar"
	"github.com/gin-gonic/gin"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"strconv"
	"time"
)

var (
	totalRequestsReceived             *expvar.Int
	totalResponsesSent                *expvar.Int
	totalProcessingTimeInMicroseconds *expvar.Int
	responseStatusCodeCount           *expvar.Map
)

func init() {
	// register expvar instances only once on package initialization
	totalRequestsReceived = expvar.NewInt(common.MetricTotalRequestsReceived)
	totalResponsesSent = expvar.NewInt(common.MetricTotalResponsesSent)
	totalProcessingTimeInMicroseconds = expvar.NewInt(common.MetricTotalProcessingTimeInMicroseconds)
	responseStatusCodeCount = expvar.NewMap(common.MetricResponseStatusCodeCount)
}

// Metrics returns useful data for analytics and debugging.
func Metrics() gin.HandlerFunc {
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
		responseStatusCodeCount.Add(strconv.Itoa(ctx.Writer.Status()), 1)
	}
}
