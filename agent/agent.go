package agent

import "time"

type Agent interface {
	RecordGauge(name string, gauge int64)
	RecordDuration(name string, duration time.Duration)
	Run()
}
