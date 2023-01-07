package structure

import (
	"sync"
	"time"
)

type (
	// Revision 代表一个区块同步过程中的某一个阶段结束，
	// 同一个 Snapshot 期间，每一个过程结束时输出一次，
	Revision struct {
		sync.Mutex

		Tag             Tag       `json:"tag"`
		Context         string    `json:"context"` // Revision 生效时设置
		Timestamp       Timestamp `json:"timestamp"`
		CommitTimestamp Timestamp `json:"commit_timestamp"` // Revision 生效时间，结合 Timestamp 和 CommitTimestamp 确定 Revision 持续时间

		SnapshotID string `json:"snapshot_id"`
	}
)

func NewRevision(tag Tag, snapshotID string) *Revision {
	return &Revision{
		Tag:        tag,
		Context:    "",
		Timestamp:  FromNow(),
		SnapshotID: snapshotID,
	}
}

// Commit 当前 Revision 生效，需要切换到下一个 Revision
func (r *Revision) Commit(context string, commitTime time.Time) error {
	r.Lock()
	defer r.Unlock()

	r.Context = context
	r.CommitTimestamp = FromTime(commitTime)
	return nil
}
