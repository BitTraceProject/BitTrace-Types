package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
)

func TestStructure(t *testing.T) {
	now := time.Now()
	initSnapshot := structure.NewInitSnapshot("1", 0, now, nil)
	finalSnapshot := initSnapshot.Commit(now.Add(time.Second*3), nil)
	t.Log(finalSnapshot)
}

func TestEncodeStructure(t *testing.T) {
	now := time.Now()
	initSnapshot := structure.NewInitSnapshot("1", 0, now, nil)
	finalSnapshot := initSnapshot.Commit(now.Add(time.Second*3), nil)
	t.Log(*finalSnapshot)

	data, err := json.Marshal(finalSnapshot)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	var dSnapshot structure.Snapshot
	err = json.Unmarshal(data, &dSnapshot)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dSnapshot)
}

func TestTime(t *testing.T) {
	now := time.Now()
	nowTs := common.FromTime(now)
	nowT := nowTs.FormatTime()
	t.Log(now, nowTs, nowT)
}
