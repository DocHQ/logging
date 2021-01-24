package logging

import (
	"testing"

	"github.com/DocHQ/logging/test"
)

func init() {
	Logger = &test.Logger{}
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

// Must run last as this os.Exits(1)
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

// Must run last as this os.Exits(1)
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

// Must run last as this os.Exits(1)
func TestDebugFatalOutput(t *testing.T) {
	Verbose = true
	Fatal("Testing fatal")
}

func TestStruct(t *testing.T) {
	Info(struct {
		Hello string
		World string
	}{"Hello", "World"})
}
