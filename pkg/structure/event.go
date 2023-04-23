package structure

import "github.com/BitTraceProject/BitTrace-Types/pkg/common"

type (
	EventOrphan struct {
		SnapshotID            string           `json:"snapshot_id"`
		Type                  EventType        `json:"event_type_orphan"`
		OrphanParentBlockHash string           `json:"orphan_parent_block_hash"`
		OrphanBlockHash       string           `json:"orphan_block_hash"`
		ConnectMainChain      bool             `json:"connect_main_chain"`
		Timestamp             common.Timestamp `json:"timestamp"` // Revision 开始时间戳
	}
	EventType int
)

const (
	EventTypeOrphanHappen EventType = iota
	EventTypeOrphanConnect
	EventTypeOrphanDiscard
	EventTypeUnknown
)

func NewEventOrphan(t EventType, snapshotID string, parentBlockHash, blockHash string, connectMainChain bool) *EventOrphan {
	e := &EventOrphan{
		SnapshotID:            snapshotID,
		Type:                  t,
		OrphanParentBlockHash: parentBlockHash,
		OrphanBlockHash:       blockHash,
		ConnectMainChain:      connectMainChain,
		Timestamp:             common.FromNow(),
	}
	return e
}
