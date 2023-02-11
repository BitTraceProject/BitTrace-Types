package config

type (
	ResolverMgrConfig struct {
		MetaServerAddr string         `mapstructure:"meta_server_addr"` // meta 服务地址
		ResolverConfig ResolverConfig `mapstructure:"resolver_config"`  // resolver 的配置
	}
	ResolverConfig struct {
		MetaServerAddr string         `mapstructure:"meta_server_addr"` // meta 服务地址
		MqServerAddr   string         `mapstructure:"mq_server_addr"`   // mq 服务地址
		DatabaseConfig DatabaseConfig `mapstructure:"database_config"`  // master
	}
)

func (conf *ResolverMgrConfig) Validate() bool {
	return conf.MetaServerAddr != "" && conf.ResolverConfig.Validate()
}

func (conf *ResolverMgrConfig) Complete() {
}

func (conf *ResolverConfig) Validate() bool {
	return conf.MetaServerAddr != "" && conf.MqServerAddr != "" && conf.DatabaseConfig.Validate()
}

func (conf *ResolverConfig) Complete() {
}
