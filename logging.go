package logging

import (
	"fmt"

	"github.com/DocHQ/logging/console"
)

var (
	Verbose bool            = false
	Logger  LoggerInterface = &console.Logger{}
)

type Level uint32

const (
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DEBUG = 0
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	INFO = 1
	// WarnLevel level. Non-critical entries that deserve eyes.
	WARN = 2
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR = 3
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FATAL = 4
)

// Some new stuff
// The default text logger

// The logger interface implements a single function to handle the desired response
// for the output
type LoggerInterface interface {
	// Log(data, key=value fields, log level, Verbose (Debug))
	Log(interface{}, map[string]interface{}, uint32, bool)
}

// The useual defines
func Error(err interface{}) {
	Logger.Log(err, make(map[string]interface{}, 0), ERROR, Verbose)
}
func Debug(debug interface{}) {
	Logger.Log(debug, make(map[string]interface{}, 0), DEBUG, Verbose)
}
func Info(info interface{}) {
	Logger.Log(info, make(map[string]interface{}, 0), INFO, Verbose)
}
func Warn(warn interface{}) {
	Logger.Log(warn, make(map[string]interface{}, 0), WARN, Verbose)
}
func Fatal(fatal interface{}) {
	Logger.Log(fatal, make(map[string]interface{}, 0), FATAL, Verbose)
}

func Errorf(form string, vars ...interface{}) {
	Logger.Log(fmt.Sprintf(form, vars...), make(map[string]interface{}, 0), ERROR, Verbose)
}
func Debugf(form string, vars ...interface{}) {
	Logger.Log(fmt.Sprintf(form, vars...), make(map[string]interface{}, 0), DEBUG, Verbose)
}
func Infof(form string, vars ...interface{}) {
	Logger.Log(fmt.Sprintf(form, vars...), make(map[string]interface{}, 0), INFO, Verbose)
}
func Warnf(form string, vars ...interface{}) {
	Logger.Log(fmt.Sprintf(form, vars...), make(map[string]interface{}, 0), WARN, Verbose)
}
func Fatalf(form string, vars ...interface{}) {
	Logger.Log(fmt.Sprintf(form, vars...), make(map[string]interface{}, 0), FATAL, Verbose)
}
