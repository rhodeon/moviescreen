package common

// expvar metric keys
const (
	MetricVersion                           = "version"
	MetricTimestamp                         = "timestamp"
	MetricGoroutines                        = "goroutines"
	MetricDatabase                          = "database"
	MetricTotalRequestsReceived             = "total_requests_received"
	MetricTotalResponsesSent                = "total_responses_sent"
	MetricTotalProcessingTimeInMicroseconds = "total_processing_time_microseconds"
	MetricResponseStatusCodeCount           = "response_status_code_count"
)
