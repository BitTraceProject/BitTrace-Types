package config

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/errorx"

	"github.com/spf13/viper"
)

type Config interface {
	Validate() bool
	Complete()
}

func NewExporterConfig(confPath string) (*ExporterConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName("exporter_config")
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType("yaml")
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(ExporterConfig)
	err := viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}

func NewReceiverConfig(confPath string) (*ReceiverConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName("receiver_config")
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType("yaml")
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(ReceiverConfig)
	err := viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}

func NewResolverMgrConfig(confPath string) (*ResolverMgrConfig, error) {
	var viperConfig = viper.New()
	viperConfig.SetConfigName("resolver_mgr_config")
	viperConfig.SetConfigFile(confPath)
	viperConfig.SetConfigType("yaml")
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := new(ResolverMgrConfig)
	err := viperConfig.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	conf.Complete()
	if ok := conf.Validate(); !ok {
		return nil, errorx.ErrConfigInvalid
	}
	return conf, nil
}
