package tests

import (
	"testing"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

func TestGap(t *testing.T) {
	key := common.GenExporterInfoKey("exporter_1")
	tag := common.ParseExporterTagFromExporterInfoKey(key)
	t.Log(key, tag)
}
