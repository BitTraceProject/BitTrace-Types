package metric

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
	"github.com/BitTraceProject/BitTrace-Types/pkg/logger"
)

type (
	MetricSnapshotStage struct {
		Tag        string    `json:"tag"`
		SnapshotID string    `json:"snapshot_id"`
		Timestamp  string    `json:"timestamp"`
		Stage      StageType `json:"snapshot_stage"`
	}
	MetricModuleError struct {
		Tag       string     `json:"tag"`
		Timestamp string     `json:"timestamp"`
		Module    ModuleType `json:"snapshot_stage"`
	}
	ModuleType string
	StageType  string
)

const (
	ModuleTypeExporter        ModuleType = "exporter"
	ModuleTypeReceiver        ModuleType = "receiver"
	ModuleTypeMq              ModuleType = "mq"
	ModuleTypeMeta            ModuleType = "meta"
	ModuleTypeResolver        ModuleType = "resolver"
	ModuleTypeResolverMgr     ModuleType = "resolver_mgr"
	ModuleTypeResolverHandler ModuleType = "resolver_handler"
	ModuleTypeUnknown         ModuleType = "unknown"
)

const (
	StageTypeInit    StageType = "init"
	StageTypeFinal   StageType = "final"
	StageTypePair    StageType = "pair"
	StageTypeError   StageType = "error"
	StageTypeUnknown StageType = "unknown"
)

var (
	metricSnapshotStageLogger logger.Logger
	metricModuleErrorLogger   logger.Logger
)

func init() {
	metricSnapshotStageLogger = logger.GetLogger("bittrace_metric_snapshot_stage")
	metricModuleErrorLogger = logger.GetLogger("bittrace_metric_module_error")
}

func (t ModuleType) String() string {
	return string(t)
}

func (t StageType) String() string {
	return string(t)
}

// MetricLogSnapshotStage 用于统计 snapshot 的各种时延
func MetricLogSnapshotStage(metrics ...MetricSnapshotStage) {
	for _, o := range metrics {
		metricSnapshotStageLogger.Msg(common.StructToJsonStr(o))
	}
}

// MetricLogModuleError 用于统计各个组件发生 error 的数量
func MetricLogModuleError(metrics ...MetricModuleError) {
	for _, o := range metrics {
		metricModuleErrorLogger.Msg(common.StructToJsonStr(o))
	}
}
