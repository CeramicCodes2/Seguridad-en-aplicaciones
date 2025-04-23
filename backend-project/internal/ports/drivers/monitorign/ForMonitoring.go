package ports

type ForMonitoring interface {
	SendLoggingData(data string) error
	SendAnomalyDetection(data string) error
	SendPerformanceMetrics(data string) error
	SendAlertNotification(data string) error
	SendUserActivityTracking(data string) error
}
