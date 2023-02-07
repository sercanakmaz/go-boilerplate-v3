package rabbitmqv2

import "time"

type Config struct {
	Host           string
	Port           string
	VHost          string
	Username       string
	Password       string
	ConnectionName string
	Reconnect      struct {
		MaxAttempt int
		Interval   time.Duration
	}
}
