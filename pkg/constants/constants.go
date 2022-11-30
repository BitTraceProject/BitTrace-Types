package constants

import "time"

const (
	RETRY_COUNT = 3 // 调用 http 或者 rpc 接口的重试次数

	LOG_EOF = "<EOF>" // 通过 EOF 符号标识一个日志文件的终结，方便 exporter 监听及时切换到下一个日志文件

	CONSUME_MQ_INTERVAL = 3 * time.Second // TODO 这里按照实际情况做一个估算，当 mq 被消费空之后，resolver 轮询 mq 的时间间隔

	TIME_LAYOUT_DAY = "2006-01-02" // day 时间格式

	ID_SEP_SYMBOL = "-" // id 生成的分隔符

	// 框架各个根目录
	BITTRACE_ROOT_DIR   = ".bittrace" // 部署的 bittrace 相关组件根目录，$HOME/$BITTRACE_ROOT_DIR
	BITTRACE_CLIENT_DIR = "btcd"      // btcd client 相关文件根目录，$HOME/$BITTRACE_ROOT_DIR/$BITTRACE_CLIENT_DIR
	BITTRACE_LOG_DIR    = "logfiles"  // 日志文件根目录，$HOME/$BITTRACE_ROOT_DIR/$BITTRACE_LOG_DIR
	BITTRACE_CONFIG_DIR = "configs"   // 部署的 bittrace 相关组件配置目录，$HOME/$BITTRACE_ROOT_DIR/$BITTRACE_CONFIG_DIR

	// 各组件配置文件名字
	BITTRACE_EXPORTER_CONFIG_NAME = "exporter_config.yaml"
	BITTRACE_RECEIVER_CONFIG_NAME = "receiver_config.yaml"
	BITTRACE_RESOLVER_CONFIG_NAME = "resolver_config.yaml"
)
