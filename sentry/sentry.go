package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

var levelToSentryLevel = map[uint32]sentry.Level{
	0: sentry.LevelDebug,
	1: sentry.LevelInfo,
	2: sentry.LevelWarning,
	3: sentry.LevelError,
	4: sentry.LevelFatal,
}

// Passthrough type for https://pkg.go.dev/github.com/getsentry/sentry-go#ClientOptions
type ConfigOptions sentry.ClientOptions

// This is not a usual init as the user may need to change options before hand
func InitSentry(config *ConfigOptions) error {
	sentryOptions := sentry.ClientOptions(*config)
	return sentry.Init(sentryOptions)
}

func NewConfig() *sentry.ClientOptions {
	config := &sentry.ClientOptions{}

	return config
}

func Flush(t time.Duration) {
	sentry.Flush(t)
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
