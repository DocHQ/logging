package logging

import (
	"testing"
)

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
	FatalNoExit("Testing fatal")
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
	FatalNoExit("Testing fatal")
}

// Testing debug

func TestDebugInfoOutput(t *testing.T) {
	DebugEnabled = true
	Info("Testing info")
}

func TestDebugDebugOutput(t *testing.T) {
	DebugEnabled = true
	Debug("Testing info")
}
func TestDebugWarnOutput(t *testing.T) {
	DebugEnabled = true
	Warn("Testing warn")
}
func TestDebugErrorOutput(t *testing.T) {
	DebugEnabled = true
	Error("Testing error")
}

// Must run last as this os.Exits(1)
func TestDebugFatalOutput(t *testing.T) {
	DebugEnabled = true
	FatalNoExit("Testing fatal")
}

func TestStruct(t *testing.T) {
	Info(struct {
		Hello string
		World string
	}{"Hello", "World"})
}

func TestInfoOutputWithTime(t *testing.T) {
	TimeEnabled = true
	Info("Testing info")
}
