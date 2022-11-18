package protocol

type (
	Message struct {
		Tag string `json:"tag"`
		Msg []byte `json:"msg"` // data+seq
	}

	PushMessageArgs struct {
		Message Message `json:"message"`
	}
	PushMessageReply struct {
		OK bool `json:"ok"`
	}

	FilterMessageArgs struct {
		Tag string `json:"tag"`
	}
	FilterMessageReply struct {
		Message Message `json:"message"`
		HasNext bool    `json:"has_next"` // 判断是否有后续消息待消费
		OK      bool    `json:"ok"`
	}
)
