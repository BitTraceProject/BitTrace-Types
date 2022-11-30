package protocol

// receiver 目前只需要一个接口, 即接收来自任意 Exporter 的快照

type (
	ReceiveData struct {
		Day  int64  `json:"day"`  // 当前传送快照的日期，防止乱序
		Seq  int64  `json:"seq"`  // 当前传送快照的序号，防止乱序
		Data []byte `json:"data"` // Revision 序列化的 Data 主体
	}
	// JoinRequest exporter 加入，开启采集任务
	JoinRequest struct {
		ExporterTag string `json:"exporter_tag"` // 按照顺序检查该 tag，检查通过分配 resolver，失败返回错误信息
	}
	// JoinResponse 返回加入结果和相关消息
	JoinResponse struct {
		OK  bool  `json:"ok"`  // 是否加入成功
		Msg int64 `json:"msg"` // 开启成功时为 resolver id，失败时为错误信息，未存在、已存在或者已删除
	}
	// ReceiveRequest 来自 Exporter, 用于上报原始快照数据
	ReceiveRequest struct {
		ExporterTag string      `json:"exporter_tag"` // 标识当前接收到的数据来自哪一个 exporter
		Data        ReceiveData `json:"data"`
	}
	// ReceiveResponse 响应，返回下一个需要的 Revision 的序号，
	// TODO 这里的返回值暂时用不到，暂时不考虑丢数据的情况（exporter 和 receiver 互相配合完成重传，累计确认，保证不丢数据）
	ReceiveResponse struct {
		NextDay int64 `json:"next_day"` // 下一个需要 Revision 的序号
		NextSeq int64 `json:"next_seq"` // 下一个需要 Revision 的序号，（当 next seq 超出文件中实际的 seq 时，说明已经切换到下一个文件了，这时候 exporter 会忽略它）
	}
	// QuitRequest exporter 退出，中止采集任务
	QuitRequest struct {
		ExporterTag string `json:"exporter_tag"` // 按照顺序检查该 tag，检查通过分配 resolver，失败返回错误信息
		LazyQuit    bool   `json:"lazy_quit"`    // 中止 resolver 时是否等处理完所有消息，或者直接终止清空该 tag 的消息，默认 false
	}
	// QuitResponse 返回退出结果和相关消息
	QuitResponse struct {
		OK  bool  `json:"ok"`  // 是否退出成功
		Msg int64 `json:"msg"` // 中止成功时为 resolver id，失败时为错误信息，未存在、已存在或者已删除
	}
)
