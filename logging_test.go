package logging

import (
	"testing"

	"github.com/DocHQ/logging/test"
)

var testData = map[string]interface{}{
	"hello": "World",
}

func init() {
	Logger = []LoggerInterface{&test.Logger{}}
}

func TestInfoOutput(t *testing.T) {
	Info("Testing info")
}
func TestDebugOutput(t *testing.T) {
	Debug("Testing info")
}
func TestWarnOutput(t *testing.T) {
	Warn("Testing warn")
}
func TestErrorOutput(t *testing.T) {
	Error("Testing error")
}
func TestFatalOutput(t *testing.T) {
	Fatal("Testing fatal")
}

// Testing veradics
func TestInfoVOutput(t *testing.T) {
	Infof("%v", "Testing info")
}
func TestDebugVOutput(t *testing.T) {
	Debugf("%v", "Testing info")
}
func TestWarnVOutput(t *testing.T) {
	Warnf("%+v", "Testing warn")
}
func TestErrorVOutput(t *testing.T) {
	Errorf("%+v", "Testing error")
}
func TestFatalVOutput(t *testing.T) {
	Fatal("Testing fatal")
}

// Testing debug

func TestDebugInfoOutput(t *testing.T) {
	Verbose = true
	Info("Testing info")
}

func TestDebugDebugOutput(t *testing.T) {
	Verbose = true
	Debug("Testing info")
}
func TestDebugWarnOutput(t *testing.T) {
	Verbose = true
	Warn("Testing warn")
}
func TestDebugErrorOutput(t *testing.T) {
	Verbose = true
	Error("Testing error")
}
func TestDebugFatalOutput(t *testing.T) {
	Verbose = true
	Fatal("Testing fatal")
}

//Testing data
func TestInfoDataOutput(t *testing.T) {
	InfoWithData("Testing info", testData)
}
func TestDebugDataOutput(t *testing.T) {
	DebugWithData("Testing info", testData)
}
func TestWarnDataOutput(t *testing.T) {
	WarnWithData("Testing warn", testData)
}
func TestErrorDataOutput(t *testing.T) {
	ErrorWithData("Testing error", testData)
}
func TestFatalDataOutput(t *testing.T) {
	FatalWithData("Testing fatal", testData)
}

func TestStruct(t *testing.T) {
	Info(struct {
		Hello string
		World string
	}{"Hello", "World"})
}
