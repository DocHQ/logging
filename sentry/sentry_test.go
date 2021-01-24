package sentry

import (
	"time"
)

func ExampleInitSentry() {
	if err := InitSentry(&ConfigOptions{
		Dsn:              "",
		Debug:            true,
		AttachStacktrace: true,
	}); err != nil {
		return
	}
	defer Flush(5 * time.Second)
	// Logger = []LoggerInterface{&test.Logger{}, &sentry.Logger{}}
}
