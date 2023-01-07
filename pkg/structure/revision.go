package structure

import (
	"sync"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/errorx"
)

type (
	// Revision 代表一个区块同步过程中的某一个阶段结束，
	// 同一个 Snapshot 期间，每一个过程结束时输出一次，
	// Revision 将多个 Status 迁移，Event，Result 等打包输出
	// TODO 并发支持改造
	Revision struct {
		sync.Mutex

		Tag             Tag       `json:"tag"`
		Context         string    `json:"context"` // Revision 生效时设置
		Timestamp       Timestamp `json:"timestamp"`
		CommitTimestamp Timestamp `json:"commit_timestamp"` // Revision 生效时间，结合 Timestamp 和 CommitTimestamp 确定 Revision 持续时间

		SnapshotID string `json:"snapshot_id"`

		// EventQueue 保存当前 Revision 期间未出结果的 Event
		// chainID + eventTag 可以唯一标识一个 Event，不采用 snapshotID 来标识是因为后续可能 snapshotID 获取困难（height 已经改变，initTimestamp 也无法得知）
		// chainID 也不能去掉，后续可能有多个 revision 并发进行，需要能够识别出 event 属于哪一个 revision，ReceiveEvent 和 CommitEvent 都会对比 chainID 是否匹配
		EventQueue map[string]Event `json:"event_queue"` // 还没有 result 发生的临时状态迁移，这一部分不被打包，Commit 时会清空

		// 逻辑关系，不会维护
		// StatusTransferList []*StatusTransfer `json:"status_transfer_list"` // result 已经发生了的最终状态迁移，revision 处不再维护 st，会实时的通过 logger 输出
	}
)

func NewRevision(tag Tag, snapshotID string) *Revision {
	return &Revision{
		Tag:        tag,
		Context:    "",
		Timestamp:  FromNow(),
		SnapshotID: snapshotID,
		EventQueue: map[string]Event{},
		//StatusTransferList: []*StatusTransfer{},
	}
}

// ReceiveEvent 此时 event 对应的 result 还未发生，event 到达了 revision，放入 map
func (r *Revision) ReceiveEvent(chainID string, event Event) bool {
	r.Lock()
	defer r.Unlock()

	if chainID != ParseChainIDFromSnapshotID(r.SnapshotID) {
		return false
	}
	eventID := GenEventID(chainID, event.Tag.String())
	_, ok := r.EventQueue[eventID]
	if ok {
		return false
	}
	r.EventQueue[eventID] = event
	return true
}

// CommitStatusTransfer 某一 event 对应的 result 已经发生了后调用的，将 event 与 result 关联起来，然后将当前的 transfer 添加到 transfer list
func (r *Revision) CommitStatusTransfer(trans *StatusTransfer, eventTag Tag) *StatusTransfer {
	r.Lock()
	defer r.Unlock()

	if trans.ChainID != ParseChainIDFromSnapshotID(r.SnapshotID) {
		return trans
	}
	eventID := GenEventID(trans.ChainID, eventTag.String())
	event, ok := r.EventQueue[eventID]
	if !ok {
		return trans
	}
	trans.RelevantEvent = event
	//r.StatusTransferList = append(r.StatusTransferList, trans)
	delete(r.EventQueue, eventID)
	return trans
}

// Commit 当前 Revision 生效，这里会把 eventqueue 清空，因为 eventqueue 是不必在 revision 中维护的，需要切换到下一个 Revision
func (r *Revision) Commit(context string, commitTime time.Time) error {
	r.Lock()
	defer r.Unlock()

	// 判断 event 是否已经清空
	if len(r.EventQueue) > 0 {
		// 未清空返回错误
		return errorx.ErrRevisionNotCommit
	}
	r.Context = context
	r.CommitTimestamp = FromTime(commitTime)
	return nil
}
