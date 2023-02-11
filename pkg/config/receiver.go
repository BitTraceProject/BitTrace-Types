package config

type ReceiverConfig struct {
	MetaServerAddr        string         `mapstructure:"meta_server_addr"`         // meta 服务地址，各种元信息以 kv 的方式由 meta 服务管理
	MqServerAddr          string         `mapstructure:"mq_server_addr"`           // receiver 服务在接收到消息后将消息放到 mq 中
	ResolverMgrServerAddr string         `mapstructure:"resolver_mgr_server_addr"` // resolver mgr 提供 resolver 的调度接口
	DatabaseConfig        DatabaseConfig `mapstructure:"database_config"`          // master
}

func (conf *ReceiverConfig) Validate() bool {
	return conf.MetaServerAddr != "" && conf.MqServerAddr != "" && conf.ResolverMgrServerAddr != "" && conf.DatabaseConfig.Validate()
}

func (conf *ReceiverConfig) Complete() {
}
