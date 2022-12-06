package constants

import "time"

// TODO 这里很多常量都要按照实际测试情况做一个估算
const (
	RETRY_COUNT = 3 // 调用 http 或者 rpc 接口的重试次数

	LOG_EOF_DAY       = "<EOF_DAY>" // 通过 EOF 符号标识一个日志文件当天的终结，方便 exporter 监听及时切换到下一个日志文件
	LOG_TEMP_FILENAME = "temp.log"  // 临时文件的文件名
	LOG_MAX_LINE      = 10000       // exporter 按数据文件来读取，超出每个数据文件的最大条目数会切换到下一个文件

	EXPORTER_POLL_DEFAULT_INTERVAL       = 100         // 100ms，exporter 定时检查是否新数据文件的间隔，该间隔内的数据会被打包，如果超出最大值，会先拆分再多打包几份
	EXPORTER_POLL_DEFAULT_INTERVAL_BLOCK = 1000        // 1000ms，exporter 定时检查是否新数据文件的间隔，该间隔内的数据会被打包，如果超出最大值，会先拆分再多打包几份
	RECEIVE_DATA_PACKAGE_MAXSIZE         = 1000 * 1000 // 1000KBi，将 exporter 的数据打包，不得超过此最大值，暂时没用到
	RECEIVE_DATA_PACKAGE_MAXN            = 10 * 1000   // 10KBi条 snapshot，将 exporter 的数据打包，不得超过此最大值

	MQ_CONSUME_INTERVAL = 300 * time.Millisecond // 300ms

	TIME_LAYOUT_DAY = "2006-01-02" // day 时间格式

	LOG_FILENAME_SEP_SYMBOL = "_" // log filename 生成的分隔符
	ID_SEP_SYMBOL           = "-" // id 生成的分隔符

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
