package tests

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
	"testing"
	"time"
)

func TestStructure(t *testing.T) {
	now := time.Now()
	initSnapshot := structure.InitSnapshot("1", 0, now, structure.NewStatus(nil, nil))
	finalSnapshot := structure.FinalSnapshot(initSnapshot, now.Add(time.Second*3), structure.NewStatus(nil, nil))
	t.Log(finalSnapshot)
}
