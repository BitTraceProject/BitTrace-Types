package config

type ExporterConfig struct {
	Tag                string `mapstructure:"tag"`                  // 唯一标识一个 exporter
	BasePath           string `mapstructure:"basepath"`             // exporter 监听的日志 base 文件夹路径
	ReceiverServerAddr string `mapstructure:"receiver_server_addr"` // receiver 服务地址
}
