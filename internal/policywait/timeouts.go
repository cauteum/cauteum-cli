package policywait

import "time"

// Operational defaults owned by this package.
const (
	defaultWaitTimeout  = 60 * time.Second
	defaultPollInterval = 200 * time.Millisecond
)
