package tests

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/metric"
	"testing"
)

func TestMetric(t *testing.T) {
	metricS := []metric.MetricSnapshotStage{
		{
			Tag:        "1",
			SnapshotID: "1",
			Timestamp:  "1",
			Stage:      metric.StageTypeInit,
		},
		{
			Tag:        "1",
			SnapshotID: "2",
			Timestamp:  "2",
			Stage:      metric.StageTypeFinal,
		},
		{
			Tag:        "2",
			SnapshotID: "1",
			Timestamp:  "3",
			Stage:      metric.StageTypeError,
		},
	}
	metricE := []metric.MetricModuleError{
		{
			Tag:       "1",
			Timestamp: "1",
			Module:    metric.ModuleTypeExporter,
		},
		{
			Tag:       "1",
			Timestamp: "2",
			Module:    metric.ModuleTypeReceiver,
		},
	}
	metric.MetricLogSnapshotStage(metricS...)
	metric.MetricLogModuleError(metricE...)
}
