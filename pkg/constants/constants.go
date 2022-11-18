package constants

import "time"

const (
	RETRY_COUNT = 3 // 调用 http 或者 rpc 接口的重试次数

	LOG_EOF = "<EOF>" // 通过 EOF 符号标识一个日志文件的终结，方便 exporter 监听及时切换到下一个日志文件

	CONSUME_MQ_INTERVAL = 3 * time.Second // TODO 这里按照实际情况做一个估算，当 mq 被消费空之后，resolver 轮询 mq 的时间间隔

	TIME_LAYOUT_DAY = "2006-01-02" // day 时间格式

	ID_SEP_SYMBOL = "-" // id 生成的分隔符
)
