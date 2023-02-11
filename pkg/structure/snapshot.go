package structure

import (
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

type (
	// Snapshot 代表一次完整的区块同步流程，主要由一系列的 Revision 组成，
	// 输出两次，初始化一次，结束一次，之间所有 ID 一致的 Revision 都属于这个 Snapshot，
	// 在 Resolver 处会将一对 Snapshot 合并为一个，然后写入数据库
	Snapshot struct {
		ID                string           `json:"snapshot_id"` // 19+1+2+1+10：init timestamp string + chain id + chain height
		TargetChainID     string           `json:"target_chain_id"`
		TargetChainHeight int32            `json:"target_chain_height"`
		Type              SnapshotType     `json:"snapshot_type"`
		Timestamp         common.Timestamp `json:"timestamp"`

		// for init and final snapshot
		RevisionList []*Revision `json:"revision_list"`

		// for all snapshot
		State *BestState `json:"state"` // 当前先用 any，目前不包括任何信息，后面可以加
	}
	SnapshotType int
)

const (
	SnapshotTypeInit = iota
	SnapshotTypeFinal
	SnapshotTypeSync
	SnapshotTypeUnknown
)

func NewInitSnapshot(targetChainID string, targetChainHeight int32, initTime time.Time, state *BestState) *Snapshot {
	ts := common.FromTime(initTime)
	id := common.GenSnapshotID(targetChainID, targetChainHeight, ts)
	s := &Snapshot{
		ID:                id,
		TargetChainID:     targetChainID,
		TargetChainHeight: targetChainHeight,
		Type:              SnapshotTypeInit,
		Timestamp:         ts,

		RevisionList: []*Revision{},

		State: state,
	}
	return s
}

func NewSyncSnapshot(bestChainID string, bestChainHeight int32, syncTime time.Time, state *BestState) *Snapshot {
	ts := common.FromTime(syncTime)
	id := common.GenSnapshotID(bestChainID, bestChainHeight, ts)
	s := &Snapshot{
		ID:                id,
		TargetChainID:     bestChainID,
		TargetChainHeight: bestChainHeight,
		Type:              SnapshotTypeSync,
		Timestamp:         ts,

		State: state,
	}
	return s
}

// CommitRevision 被 init snapshot 调用
func (s *Snapshot) CommitRevision(revision *Revision) {
	s.RevisionList = append(s.RevisionList, revision)
}

// Commit 被 init snapshot 调用
func (s *Snapshot) Commit(finalTime time.Time, state *BestState) *Snapshot {
	ts := common.FromTime(finalTime)
	return &Snapshot{
		ID:                s.ID,
		TargetChainID:     s.TargetChainID,
		TargetChainHeight: s.TargetChainHeight,
		Type:              SnapshotTypeFinal,
		Timestamp:         ts,
		RevisionList:      s.RevisionList,
		State:             state,
	}
}
