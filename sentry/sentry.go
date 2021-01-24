package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

var (
	Config = sentry.ClientOptions{}
)

var levelToSentryLevel = map[uint32]sentry.Level{
	0: sentry.LevelDebug,
	1: sentry.LevelInfo,
	2: sentry.LevelWarning,
	3: sentry.LevelError,
	4: sentry.LevelFatal,
}

// This is not a usual init as the user may need to change options before hand
func InitSentry() error {
	// Set this up so you dont need to call flush() each log
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 3
	Config.Transport = sentrySyncTransport

	return sentry.Init(Config)
}

type Logger struct{}

func (t Logger) Log(i interface{}, fields map[string]interface{}, level uint32, verbose bool) {
	localHub := sentry.CurrentHub().Clone()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		// If extra fields were provided, use them
		if len(fields) != 0 {
			for k, v := range fields {
				scope.SetExtra(k, v)
			}
		}
		scope.SetLevel(levelToSentryLevel[level])
	})

	if level == 0 {
		return // Don't log debug messages
	}

	switch msg := i.(type) {
	case error:
		localHub.CaptureException(msg)
	case string:
		localHub.CaptureMessage(msg)
	default:
		localHub.CaptureMessage(fmt.Sprint(msg))
	}
}
