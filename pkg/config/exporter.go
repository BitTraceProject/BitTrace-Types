package config

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
)

type (
	ExporterConfig struct {
		ExporterTag string `mapstructure:"exporter_tag"` // 作为 exporter 当前的唯一 Tag（如果唯一性验证失败会 panic），同时通过路径组装生成日志文件路径扫描日志文件

		ReceiverServerAddr string `mapstructure:"receiver_server_addr"` // receiver 服务地址

		PollInterval int `mapstructure:"poll_interval"` // 自定义 export 时间间隔，单位为 1ms，默认为 100ms，不宜设置过小

		FileKeepingDays int `mapstructure:"file_keeping_days"` // 已上报的日志，保留多久后会被删除
	}
)

func (conf *ExporterConfig) Validate() bool {
	// TODO 这里暂不检验 day 格式和 seq 数值
	return conf != nil && conf.ExporterTag != "" && len(conf.ExporterTag) <= constants.EXPORTER_MAX_TAG_LENGTH && conf.ReceiverServerAddr != ""
}

func (conf *ExporterConfig) Complete() {
	if conf.PollInterval == 0 {
		conf.PollInterval = constants.EXPORTER_POLL_DEFAULT_INTERVAL
	}
}
