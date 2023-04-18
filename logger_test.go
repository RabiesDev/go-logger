package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logger := NewLogger(os.Stdout, DefaultPrefixes(), 0).WithColor()
	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
	logger.Fatal("Fatal")
}
