package protocol

// receiver 目前只需要一个接口, 即接收来自任意 Exporter 的快照

type (
	// ReceiverRequest 来自 Exporter, 用于上报原始快照数据
	ReceiverRequest struct {
		Tag  string `json:"tag"`  // 标识当前接收到的数据来自哪一个 exporter
		Data []byte `json:"data"` // Revision 序列化的 Data 主体
		Seq  int64  `json:"seq"`  // 当前传送快照的序号，防止乱序
	}
	// ReceiverResponse 响应，返回下一个需要的 Revision 的序号
	ReceiverResponse struct {
		Seq int64 `json:"seq"` // 下一个需要 Revision 的序号
	}
)
