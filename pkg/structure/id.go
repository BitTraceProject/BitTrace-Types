package structure

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"strconv"
	"strings"
)

func GenSnapshotID(targetChainID string, targetChainHeight int32, initTimestamp Timestamp) string {
	return strings.Join([]string{
		targetChainID,
		strconv.FormatInt(int64(targetChainHeight), 10),
		initTimestamp.Format(constants.TIME_LAYOUT_DAY)}, constants.ID_SEP_SYMBOL)
}

func GenEventID(chainID string, eventTag string) string {
	return strings.Join([]string{chainID, eventTag}, constants.ID_SEP_SYMBOL)
}

// GenChainID 根据当前链所处的 fork 的区块高度作为该 chain 的唯一标识，
// TODO 是否存在不同 fork 高度相同的情况
func GenChainID(forkHeight int32) string {
	return strconv.FormatInt(int64(forkHeight), 10)
}

func ParseChainIDFromSnapshotID(snapshotID string) string {
	segs := strings.Split(snapshotID, constants.ID_SEP_SYMBOL)
	if len(segs) != 3 {
		return ""
	}
	return segs[0]
}
