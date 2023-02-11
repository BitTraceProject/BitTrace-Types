package protocol

// receiver 目前只需要一个接口, 即接收来自任意 Exporter 的快照

type (
	ReceiverDataPackage struct {
		Day         string   `json:"day"`          // 当前传送数据包裹的日期，用于排序
		LeftSeq     int64    `json:"left_seq"`     // 当前传送数据包裹的左序号，用于排序
		RightSeq    int64    `json:"right_seq"`    // 当前传送数据包裹的右序号，用于排序
		DataPackage [][]byte `json:"data_package"` // 快照序列化的 Data Package 主体 <=> Snapshot list
		EOF         bool     `json:"eof"`
	}

	// ReceiverJoinRequest GET exporter 加入，开启采集任务，不存在会创建
	ReceiverJoinRequest struct {
		ExporterTag string `json:"exporter_tag"` // 按照顺序检查该 tag，检查通过分配 resolver，失败返回错误信息
	}
	// ReceiverJoinResponse 返回加入结果和相关消息
	ReceiverJoinResponse struct {
		OK   bool         `json:"ok"`  // 是否加入成功
		Msg  string       `json:"msg"` // 开启成功时为 resolver id，失败时为错误信息
		Info ExporterInfo `json:"info"`
	}

	// ReceiverQueryRequest GET 查询 exporter 进度等信息，不存在不会创建
	ReceiverQueryRequest struct {
		ExporterTag string `json:"exporter_tag"` // 按照顺序检查该 tag
	}
	// ReceiverQueryResponse 返回查询结果和相关消息
	ReceiverQueryResponse struct {
		OK   bool         `json:"ok"`  // 是否查询成功
		Msg  string       `json:"msg"` // 查询成功时为 resolver id，失败时为错误信息，未存在或者已删除
		Info ExporterInfo `json:"info"`
	}

	// ReceiverDataRequest POST 来自 Exporter, 用于上报原始快照数据
	ReceiverDataRequest struct {
		ExporterTag     string              `json:"exporter_tag"` // 标识当前接收到的数据来自哪一个 exporter
		DataPackage     ReceiverDataPackage `json:"data_package"`
		CurrentProgress ExporterProgress    `json:"current_progress"`
	}
	// ReceiverDataResponse 响应，为空
	// 由于会 retry，所以必定不会丢数据（大概率）
	// TODO 需要测试，概率是多少，能否接受
	ReceiverDataResponse struct {
		OK bool `json:"ok"` // data package 是否成功放入 mq
	}

	// ReceiverQuitRequest GET exporter 退出，中止采集任务，会清空 exporter 的元信息
	ReceiverQuitRequest struct {
		ExporterTag string `json:"exporter_tag"` // 按照顺序检查该 tag，检查通过分配 resolver，失败返回错误信息
		LazyQuit    bool   `json:"lazy_quit"`    // 中止 resolver 时是否等处理完所有消息，或者直接终止清空该 tag 的消息，默认 false
	}
	// ReceiverQuitResponse 返回退出结果和相关消息
	ReceiverQuitResponse struct {
		OK  bool   `json:"ok"`  // 是否退出成功
		Msg string `json:"msg"` // 中止成功时为 resolver id，失败时为错误信息，未存在、已存在或者已删除
	}
)
