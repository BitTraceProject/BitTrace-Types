package protocol

type (
	// MqMessage MQ 中消息的结构，Tag 用于过滤消息，Msg 是消息主体
	MqMessage struct {
		Tag string `json:"tag"`
		Msg []byte `json:"msg"` // ReceiveDataPackage
	}
	// MqPushMessageArgs 添加一个消息
	MqPushMessageArgs struct {
		Message MqMessage `json:"message"`
	}
	// MqPushMessageReply 是否添加成功
	MqPushMessageReply struct {
		OK bool `json:"ok"`
	}
	// MqFilterMessageArgs 通过 Tag 过滤得到一个消息，这个消息是所有 Tag 匹配的第一个消息
	MqFilterMessageArgs struct {
		Tag string `json:"tag"`
	}
	// MqFilterMessageReply 通过 Tag 过滤得到一个消息，这个消息是所有 Tag 匹配的第一个消息
	MqFilterMessageReply struct {
		Message MqMessage `json:"message"`
		HasNext bool      `json:"has_next"` // 判断是否有后续消息待消费
		OK      bool      `json:"ok"`       // 是否过滤匹配成功
	}
	// MqClearMessageArgs 通过 Tag 过滤得到相关消息，清空并删除 tag
	MqClearMessageArgs struct {
		Tag string `json:"tag"`
	}
	// MqClearMessageReply 通过 Tag 过滤得到相关消息，清空并删除 tag
	MqClearMessageReply struct {
		Number int  `json:"number"` // 清理数目
		OK     bool `json:"ok"`     // 是否清理成功
	}
)
