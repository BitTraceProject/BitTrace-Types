package config

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
)

type (
	ExporterConfig struct {
		ReceiverServerAddr string `mapstructure:"receiver_server_addr"` // receiver 服务地址

		Tag string `mapstructure:"tag"` // 作为 exporter 当前的唯一 Tag（如果唯一性验证失败会 panic），同时通过路径组装生成日志文件路径扫描日志文件

		StartDay string `mapstructure:"start_day"` // 如 2022-12-01，根据日志文件名字过滤掉日期以前的文件
		StartSeq int64  `mapstructure:"start_seq"` // 根据 seq 信息，过滤掉该日期 seq 以前的文件

		PollInterval int64 `mapstructure:"poll_interval"` // 自定义 export 时间间隔，单位为 1ms，默认为 100ms，不宜设置过小
	}
)

func (conf *ExporterConfig) Validate() bool {
	// TODO 这里暂不检验 day 格式和 seq 数值
	return conf != nil && conf.Tag != "" && conf.ReceiverServerAddr != ""
}

func (conf *ExporterConfig) Complete() {
	if conf.PollInterval == 0 {
		conf.PollInterval = constants.EXPORTER_POLL_DEFAULT_INTERVAL
	}
}
