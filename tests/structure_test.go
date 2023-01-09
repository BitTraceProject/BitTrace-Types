package tests

import (
	"testing"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
)

func TestStructure(t *testing.T) {
	now := time.Now()
	initSnapshot := structure.NewSnapshot("1", 0, now)
	finalSnapshot := initSnapshot.Commit(now.Add(time.Second * 3))
	t.Log(finalSnapshot)
}

func TestTime(t *testing.T) {
	now := time.Now()
	nowTs := structure.FromTime(now)
	nowT := nowTs.FormatTime()
	t.Log(now, nowTs, nowT)
}
