package protocol

type (
	// Message MQ 中消息的结构，Tag 用于过滤消息，Msg 是消息主体
	Message struct {
		Tag string `json:"tag"`
		Msg []byte `json:"msg"` // data+seq
	}
	// PushMessageArgs 添加一个消息
	PushMessageArgs struct {
		Message Message `json:"message"`
	}
	// PushMessageReply 是否添加成功
	PushMessageReply struct {
		OK bool `json:"ok"`
	}
	// FilterMessageArgs 通过 Tag 过滤得到一个消息，这个消息是所有 Tag 匹配的第一个消息
	FilterMessageArgs struct {
		Tag string `json:"tag"`
	}
	// FilterMessageReply 通过 Tag 过滤得到一个消息，这个消息是所有 Tag 匹配的第一个消息
	FilterMessageReply struct {
		Message Message `json:"message"`
		HasNext bool    `json:"has_next"` // 判断是否有后续消息待消费
		OK      bool    `json:"ok"`       // 是否过滤匹配成功
	}
)
