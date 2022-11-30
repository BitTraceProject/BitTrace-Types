package config

type (
	ResolverMgrConfig struct {
		MetaServerAddr string          `mapstructure:"meta_server_addr"` // meta 服务地址
		ResolverConfig *ResolverConfig `mapstructure:"resolver_config"`  // resolver 的配置
	}
	ResolverConfig struct {
		MqServerAddr              string `mapstructure:"mq_server_addr"`               // mq 服务地址
		CollectorWriterServerAddr string `mapstructure:"collector_writer_server_addr"` // collector 服务地址
	}
)

func (conf *ResolverMgrConfig) Validate() bool {
	return conf.MetaServerAddr != "" && conf.ResolverConfig != nil && conf.ResolverConfig.Validate()
}

func (conf *ResolverMgrConfig) Complete() {
}

func (conf *ResolverConfig) Validate() bool {
	return conf.MqServerAddr != "" && conf.CollectorWriterServerAddr != ""
}

func (conf *ResolverConfig) Complete() {
}
