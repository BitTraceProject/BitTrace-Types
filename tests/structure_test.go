package tests

import (
	"testing"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
)

func TestStructure(t *testing.T) {
	now := time.Now()
	initSnapshot := structure.InitSnapshot("1", 0, now, structure.NewStatus(nil, nil))
	finalSnapshot := structure.FinalSnapshot(initSnapshot, now.Add(time.Second*3), structure.NewStatus(nil, nil))
	t.Log(finalSnapshot)
}

func TestTime(t *testing.T) {
	now := time.Now()
	nowTs := structure.FromTime(now)
	nowT := nowTs.FormatTime()
	t.Log(now, nowTs, nowT)
}
