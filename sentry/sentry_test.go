package sentry

import (
	"time"
)

func ExampleInitSentry() {
	if err := InitSentry(&ConfigOptions{
		Dsn:              "https://b28a0e6d643349e8ba8dbe671f5bbb4e@o239521.ingest.sentry.io/5605608",
		Debug:            true,
		AttachStacktrace: true,
	}); err != nil {
		return
	}
	defer Flush(5 * time.Second)
	// Logger = []LoggerInterface{&test.Logger{}, &sentry.Logger{}}
}
