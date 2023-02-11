package tests

import (
	"testing"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/logger"
)

func TestGetLogger(t *testing.T) {
	l := logger.GetLogger("test_logger")
	l.Msg("data 0")
	l.Info("test info %d", 0)
	l.Warn("test warn %d", 0)
	l.Fatal("test fatal %d", 0)
	l.Error("test error %d", 0)
	time.Sleep(3 * time.Second)
	l.Msg("data 2")
	l.Error("test error %d", 1)
	l.Error("test error %d", 2)
	l.Error("test error %d", 3)
}
