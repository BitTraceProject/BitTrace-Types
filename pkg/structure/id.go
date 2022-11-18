package structure

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"strconv"
	"strings"
)

func GenSnapshotID(targetChainID string, targetChainHeight int64, initTimestamp Timestamp) string {
	return strings.Join([]string{
		targetChainID,
		strconv.FormatInt(targetChainHeight, 10),
		initTimestamp.Format(constants.TIME_LAYOUT_DAY)}, constants.ID_SEP_SYMBOL)
}

func GenEventID(chainID string, eventTag string) string {
	return strings.Join([]string{chainID, eventTag}, constants.ID_SEP_SYMBOL)
}

func GenChainID(forkHeight int64) string {
	return strconv.FormatInt(forkHeight, 10)
}

func ParseChainIDFromSnapshotID(snapshotID string) string {
	segs := strings.Split(snapshotID, constants.ID_SEP_SYMBOL)
	if len(segs) != 3 {
		return ""
	}
	return segs[0]
}
