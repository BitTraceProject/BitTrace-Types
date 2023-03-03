package structure

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

type (
	// Snapshot 代表一次完整的区块同步流程，主要由一系列的 Revision 组成，
	// 输出两次，初始化一次，结束一次，之间所有 ID 一致的 Revision 都属于这个 Snapshot，
	// 在 Resolver 处会将一对 Snapshot 合并为一个，然后写入数据库
	Snapshot struct {
		// for all snapshot
		ID                string           `json:"snapshot_id"` // 19+1+2+1+10：init timestamp string + chain id + chain height
		TargetChainID     string           `json:"target_chain_id"`
		TargetChainHeight int32            `json:"target_chain_height"`
		Type              SnapshotType     `json:"snapshot_type"`
		Timestamp         common.Timestamp `json:"timestamp"`
		State             *BestState       `json:"state"` // 当前先用 any，目前不包括任何信息，后面可以加

		// for init and final snapshot
		BlockHash       string         `json:"block_hash"`
		IsOrphan        bool           `json:"is_orphan"`
		RevisionList    []*Revision    `json:"revision_list"`
		EventOrphanList []*EventOrphan `json:"event_orphan_list"`
	}
	SnapshotType int
)

const (
	SnapshotTypeInit = iota
	SnapshotTypeFinal
	SnapshotTypeSync
	SnapshotTypeUnknown
)

func NewInitSnapshot(targetChainID string, targetChainHeight int32, initTime time.Time, blockHash string, state *BestState) *Snapshot {
	ts := common.FromTime(initTime)
	id := common.GenSnapshotID(targetChainID, targetChainHeight, ts)
	isOrphan := false
	if targetChainID == constants.ORPHAN_CHAIN_ID {
		isOrphan = true
	}
	s := &Snapshot{
		ID:                id,
		TargetChainID:     targetChainID,
		TargetChainHeight: targetChainHeight,
		Type:              SnapshotTypeInit,
		Timestamp:         ts,
		State:             state,

		RevisionList:    []*Revision{},
		BlockHash:       blockHash,
		IsOrphan:        isOrphan,
		EventOrphanList: []*EventOrphan{},
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

// CommitOrphanEvent 被 init snapshot 调用
func (s *Snapshot) CommitOrphanEvent(eventOrphan *EventOrphan) {
	s.EventOrphanList = append(s.EventOrphanList, eventOrphan)
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
		State:             state,

		BlockHash:       s.BlockHash,
		IsOrphan:        s.IsOrphan,
		RevisionList:    s.RevisionList,
		EventOrphanList: s.EventOrphanList,
	}
}
