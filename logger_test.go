package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logger := NewLogger(os.Stdout, DefaultPrefix(), 0).WithColor()
	logger.Info("Info")
}
