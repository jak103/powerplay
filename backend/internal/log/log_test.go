package log

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestWarn(t *testing.T) {
	Warn("This is a warning %v", 1)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
