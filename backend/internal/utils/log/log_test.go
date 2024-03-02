package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init("DEBUG", false)

	assert.Equal(t, false, TheLogger.color)
	assert.Equal(t, 2, TheLogger.skip)
	assert.NotNil(t, TheLogger.tags)
}

func TestWarn(t *testing.T) {
	output := ""
	TheLogger.SetTestCapture(&output)
	Warn("This is a warning %v", 1)

	assert.Contains(t, output, "[WARN ]")
	assert.Contains(t, output, "This is a warning 1")
}

func TestDebug(t *testing.T) {
	output := ""
	TheLogger.SetTestCapture(&output)
	Debug("This is a debug %v", 1)

	assert.Contains(t, output, "[DEBUG]")
	assert.Contains(t, output, "This is a debug 1")
}

func TestInfo(t *testing.T) {
	output := ""
	TheLogger.SetTestCapture(&output)
	Info("This is a info %v", 1)

	assert.Contains(t, output, "[INFO ]")
	assert.Contains(t, output, "This is a info 1")
}

func TestError(t *testing.T) {
	output := ""
	TheLogger.SetTestCapture(&output)
	Error("This is an error %v", 1)

	assert.Contains(t, output, "[ERROR]")
	assert.Contains(t, output, "This is an error 1")
}

func TestAlert(t *testing.T) {
	output := ""
	TheLogger.SetTestCapture(&output)
	Alert("This is an alert %v", 1)

	assert.Contains(t, output, "[ALERT]")
	assert.Contains(t, output, "This is an alert 1")
}

func TestWithErr(t *testing.T) {
	Init("DEBUG", false)
	output := ""
	TheLogger.SetTestCapture(&output)
	WithErr(errors.New("test error")).Error("Testing withErr")

	assert.Contains(t, output, "[ERROR]")
	assert.Contains(t, output, "Testing withErr")
	assert.Contains(t, output, "error=test error")
}

func TestRequestId(t *testing.T) {
	Init("DEBUG", false)
	output := ""
	TheLogger.SetTestCapture(&output)
	WithRequestId("test id").Info("Testing withRequestId")

	assert.Contains(t, output, "[INFO ]")
	assert.Contains(t, output, "Testing withRequestId")
	assert.Contains(t, output, "request_id=test id")
}

func TestChainedWiths(t *testing.T) {
	Init("DEBUG", false)
	output := ""
	TheLogger.SetTestCapture(&output)
	TheLogger.WithRequestId("test id").WithErr(errors.New("new error")).Alert("Testing")

	assert.Contains(t, output, "[ALERT]")
	assert.Contains(t, output, "request_id=test id")
	assert.Contains(t, output, "error=new error")
}
