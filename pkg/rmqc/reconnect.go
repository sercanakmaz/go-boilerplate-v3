package rmqc

import "time"

type Reconnect struct {
	MaxAttempt int
	Interval   time.Duration
}
