package structure

import (
	"sync"
	"time"
)

type (
	// Snapshot 代表一次完整的区块同步流程，主要由一系列的 Revision 组成，
	// 输出两次，初始化一次，结束一次，之间所有 ID 一致的 Revision 都属于这个 Snapshot，
	// 在 Resolver 处会将一对 Snapshot 合并为一个，然后写入数据库
	Snapshot struct {
		sync.Mutex

		ID                string       `json:"id"` // init timestamp + chain id + chain height
		TargetChainID     string       `json:"target_chain_id"`
		TargetChainHeight int32        `json:"target_chain_height"`
		Type              SnapshotType `json:"type"`

		InitTimestamp  Timestamp `json:"init_timestamp"`
		FinalTimestamp Timestamp `json:"final_timestamp"`

		RevisionList []*Revision `json:"revision_list"`
	}
	SnapshotType int
)

const (
	SnapshotTypeInit = iota
	SnapshotTypeFinal
	SnapshotTypeUnknown
)

func NewSnapshot(targetChainID string, targetChainHeight int32, initTime time.Time) *Snapshot {
	timestamp := FromTime(initTime)
	id := GenSnapshotID(targetChainID, targetChainHeight, timestamp)
	s := &Snapshot{
		ID:                id,
		TargetChainID:     targetChainID,
		TargetChainHeight: targetChainHeight,
		Type:              SnapshotTypeInit,
		InitTimestamp:     timestamp,
		RevisionList:      []*Revision{},
	}
	return s
}

func (s *Snapshot) CommitRevision(revision *Revision) {
	s.Lock()
	defer s.Unlock()

	s.RevisionList = append(s.RevisionList, revision)
}

func (s *Snapshot) Commit(finalTime time.Time) *Snapshot {
	s.Lock()
	defer s.Unlock()

	timestamp := FromTime(finalTime)
	return &Snapshot{
		ID:                s.ID,
		TargetChainID:     s.TargetChainID,
		TargetChainHeight: s.TargetChainHeight,
		Type:              SnapshotTypeFinal,
		FinalTimestamp:    timestamp,
		RevisionList:      s.RevisionList,
	}
}
