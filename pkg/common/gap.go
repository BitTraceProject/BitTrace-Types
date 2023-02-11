package common

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
)

// gap: Generate and Parse

// btcd

func GenSnapshotID(targetChainID string, targetChainHeight int32, initTimestamp Timestamp) string {
	return strings.Join([]string{
		initTimestamp.String(),
		targetChainID,
		strconv.FormatInt(int64(targetChainHeight), 10),
	}, constants.ID_SEP_SYMBOL)
}

// GenChainID 根据当前链所处的 fork 的区块高度作为该 chain 的唯一标识，
// TODO 存在不同 fork 高度相同的情况，需要加以区分
func GenChainID(forkHeight int32) string {
	return strconv.FormatInt(int64(forkHeight), 10)
}

func ParseChainIDFromSnapshotID(snapshotID string) string {
	segs := strings.Split(snapshotID, constants.ID_SEP_SYMBOL)
	if len(segs) != 3 {
		return ""
	}
	return segs[1]
}

// exporter

func GenLogFilename(id int64) string {
	return strconv.FormatInt(id, 10) + ".log"
}

// GenLogFileParentPath basePath/loggerName/day/
func GenLogFileParentPath(basePath, loggerName string, day string) string {
	return filepath.Join(basePath, loggerName, day)
}

// GenLogFilepath basePath/loggerName/day/id.log
func GenLogFilepath(basePath, loggerName string, day string, id int64) string {
	return filepath.Join(GenLogFileParentPath(basePath, loggerName, day), GenLogFilename(id))
}

func CatchUpFileID(loggerName, currentDay string) (int64, int64, error) {
	var (
		currentFileID int64
		currentN      int64
	)
	basePath := GenLogFileParentPath(constants.LOGGER_FILE_BASE_PATH, loggerName, currentDay)
	ms, err := filepath.Glob(fmt.Sprintf("%s/*.log", basePath))
	if err != nil {
		return 0, 0, err
	}

	if len(ms) == 0 {
		currentFileID = 0
		currentN = 0
	} else {
		currentFileID = int64(len(ms)) - 1
		// scan file line as n
		currentFilePath := GenLogFilepath(constants.LOGGER_FILE_BASE_PATH, loggerName, currentDay, currentFileID)
		lines, _, _ := ScanFileLines(currentFilePath, 0)
		currentN = int64(len(lines))
	}
	return currentFileID, currentN, nil
}

// receiver

func GenExporterInfoKey(exporterTag string) string {
	return strings.Join([]string{exporterTag, "info"}, constants.DEFAULT_SEP_SYMBOL)
}

func ParseExporterTagFromExporterInfoKey(exporterInfoKey string) string {
	exporterTag := strings.TrimSuffix(exporterInfoKey, constants.DEFAULT_SEP_SYMBOL+"info")
	return exporterTag
}

// resolver

func GenResolverTag(exporterTag string) string {
	return strings.Join([]string{"resolver", exporterTag}, constants.DEFAULT_SEP_SYMBOL)
}

// collector

func GenSnapshotDataTableName(exporterTag string, timestamp Timestamp) string {
	return strings.Join([]string{constants.TABLE_SNAPSHOT_DATA_PREFIX, exporterTag, timestamp.String()}, constants.DEFAULT_SEP_SYMBOL)
}

func GenSnapshotSyncTableName(exporterTag string, timestamp Timestamp) string {
	return strings.Join([]string{constants.TABLE_SNAPSHOT_SYNC_PREFIX, exporterTag, timestamp.String()}, constants.DEFAULT_SEP_SYMBOL)
}

func GenStateTableName(exporterTag string, timestamp Timestamp) string {
	return strings.Join([]string{constants.TABLE_STATE_PREFIX, exporterTag, timestamp.String()}, constants.DEFAULT_SEP_SYMBOL)
}

func GenRevisionTableName(exporterTag string, timestamp Timestamp) string {
	return strings.Join([]string{constants.TABLE_REVISION_PREFIX, exporterTag, timestamp.String()}, constants.DEFAULT_SEP_SYMBOL)
}
