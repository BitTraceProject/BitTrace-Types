package config

type ReceiverConfig struct {
	MetaServerAddr        string `mapstructure:"meta_server_addr"`
	MqServerAddr          string `mapstructure:"mq_server_addr"`
	ResolverMgrServerAddr string `mapstructure:"resolver_mgr_server_addr"`
}
