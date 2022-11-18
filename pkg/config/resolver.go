package config

type (
	ResolverMgrConfig struct {
		MetaServerAddr string         `mapstructure:"meta_server_addr"`
		ResolverConfig ResolverConfig `mapstructure:"resolver_config"`
	}
	ResolverConfig struct {
		MqServerAddr        string `mapstructure:"mq_server_addr"`
		CollectorServerAddr string `mapstructure:"collector_server_addr"`
	}
)
