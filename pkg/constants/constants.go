package constants

import "time"

const (
	RETRY_COUNT = 3

	LOG_EOF = "<EOF>"

	CONSUME_MQ_INTERVAL = 3 * time.Second // TODO 这里按照实际情况做一个估算

	TIME_LAYOUT_DAY = "2006-01-02"

	ID_SEP_SYMBOL = "-"
)
