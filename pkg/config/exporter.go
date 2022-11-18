package config

type ExporterConfig struct {
	Tag                string `mapstructure:"tag"` // 唯一标识一个 exporter
	BasePath           string `mapstructure:"basepath"`
	ReceiverServerAddr string `mapstructure:"receiver_server_addr"`
}
