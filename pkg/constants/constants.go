package constants

import (
	"time"
)

// TODO 这里很多常量都要按照实际测试情况做一个估算
const (
	DEFAULT_RETRY_COUNT = 3 // 调用接口的默认重试次数，rpc、http 等

	LOGGER_EOF_DAY              = "<EOF_DAY>" // 通过 EOF 符号标识一个日志文件当天的终结，方便 exporter 监听及时切换到下一个日志文件
	LOGGER_SYNC_HEIGHT_INTERVAL = 500         // 每隔 500 左右的区块写入一次同步标记，同步较慢时，还需要其他机制，来保证每一天至少写入一条同步标记
	LOGGER_FILE_BASE_PATH       = "/root/" + BITTRACE_ROOT_DIR + "/" + BITTRACE_LOG_DIR

	ORPHAN_CHAIN_ID = "-1" // 如果当前区块同步过程中的区块被处理成区块，则将其 ChainID 设为 -1

	EXPORTER_MAX_TAG_LENGTH                 = 30      // 受 table name 最长为 64 的限制
	EXPORTER_POLL_DEFAULT_INTERVAL          = 100     // 100ms，exporter 定时检查是否新数据文件的间隔，该间隔内的数据会被打包，如果超出最大值，会先拆分再多打包几份
	EXPORTER_POLL_DEFAULT_INTERVAL_BLOCKING = 1000    // 1000ms，发生阻塞时调整的间隔
	EXPORTER_DATA_PACKAGE_MAXN              = 2 * 100 // 200条 snapshot，将数据打包，Logger 据此划分文件

	RESOLVER_CONSUME_MQ_INTERVAL      = 300 * time.Millisecond // 300ms
	RESOLVER_SNAPSHOTPAIR_STREAM_SIZE = 10                     // resolver handler 处 snapshot pair stream 的大小，注意：chan 里的元素是 snapshotPair list，不是单个数据结构

	TIME_LAYOUT_DAY = "2006-01-02" // day 时间格式

	DEFAULT_SEP_SYMBOL = "_" // tag, key, name, filepath, etc.
	ID_SEP_SYMBOL      = "-" // id 生成的分隔符

	// 框架各个根目录
	BITTRACE_ROOT_DIR   = ".bittrace" // 部署的 bittrace 相关组件根目录，/root/$BITTRACE_ROOT_DIR
	BITTRACE_CLIENT_DIR = "btcd"      // btcd client 相关文件根目录，/root/$BITTRACE_ROOT_DIR/$BITTRACE_CLIENT_DIR
	BITTRACE_LOG_DIR    = "logfiles"  // 日志文件根目录，/root/$BITTRACE_ROOT_DIR/$BITTRACE_LOG_DIR
	BITTRACE_CONFIG_DIR = "configs"   // 部署的 bittrace 相关组件配置目录，/root/$BITTRACE_ROOT_DIR/$BITTRACE_CONFIG_DIR

	// 各组件配置文件名字
	BITTRACE_EXPORTER_CONFIG_NAME = "exporter_config.yaml"
	BITTRACE_RECEIVER_CONFIG_NAME = "receiver_config.yaml"
	BITTRACE_RESOLVER_CONFIG_NAME = "resolver_config.yaml"

	// bittrace 数据库相关
	DATABASE_NAME_BITTRACE     = "bittrace"
	TABLE_SNAPSHOT_DATA_PREFIX = "snapshot_data"
	TABLE_SNAPSHOT_SYNC_PREFIX = "snapshot_sync"
	TABLE_STATE_PREFIX         = "state"
	TABLE_REVISION_PREFIX      = "revision"
	TABLE_EVENT_ORPHAN_PREFIX  = "event_orphan"

	// bittrace_openapi 数据库相关
	DATABASE_NAME_BITTRACE_OPENAPI = "bittrace_openapi"
	// ...
)
