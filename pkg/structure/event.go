package structure

import "github.com/BitTraceProject/BitTrace-Types/pkg/common"

type (
	EventOrphan struct {
		SnapshotID      string           `json:"snapshot_id"`
		Type            EventType        `json:"event_type_orphan"`
		OrphanBlockHash string           `json:"orphan_block_hash"`
		Timestamp       common.Timestamp `json:"timestamp"` // Revision 开始时间戳
	}
	EventType int
)

const (
	EventTypeOrphanHappen EventType = iota
	EventTypeOrphanConnect
	EventTypeOrphanDiscard
	EventTypeUnknown
)

func NewEventOrphan(t EventType, snapshotID string, blockHash string) *EventOrphan {
	e := &EventOrphan{
		Type:            t,
		OrphanBlockHash: blockHash,
		Timestamp:       common.FromNow(),
		SnapshotID:      snapshotID,
	}
	return e
}
