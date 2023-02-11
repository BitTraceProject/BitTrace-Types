package common

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"github.com/BitTraceProject/BitTrace-Types/pkg/errorx"
)

func ExecuteFunctionByRetry(f func() error) error {
	var err error
	for i := 0; i < constants.DEFAULT_RETRY_COUNT; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}

// ScanFileLines 从 startN 处开始，scan 数据到达文件结尾，返回数据数组，eof 和 error，
// eof 判断是按照自定义的一行来判断的，注意这里潜在的问题是 logger 与 exporter 读写竞争可能导致日志输出到错误的 day，
// 所以在 resolver 处理时，只依赖数据本身，不依赖 logger 的信息，logger 和 exporter 只需要保证不丢数据就行
func ScanFileLines(filePath string, startN int64) ([][]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	//s := bufio.NewScanner(f)
	for i := int64(0); i < startN; i++ {
		line, err := r.ReadBytes('\n')
		if err != nil {
			return nil, false, nil
		}
		// 当前 n 处还没有数据，但是已经 eof 了
		if string(line[:len(line)-1]) == constants.LOGGER_EOF_DAY {
			return nil, true, nil
		}
	}
	var lines = make([][]byte, 0)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		if string(line) == constants.LOGGER_EOF_DAY {
			return lines, true, nil
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, false, nil
}

func IsFileExisted(filePath string) bool {
	info, err := os.Stat(filePath)
	return (err == nil || os.IsExist(err)) && !info.IsDir()
}

func IsDirExisted(dirPath string) bool {
	info, err := os.Stat(dirPath)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}

func DirFileCount(dirPath string) int {
	files, _ := os.ReadDir(dirPath)
	return len(files)
}

func DayTime(dayStr string) time.Time {
	day, _ := time.Parse(constants.TIME_LAYOUT_DAY, dayStr)
	return day
}

func CurrentDay(day time.Time) string {
	return day.Format(constants.TIME_LAYOUT_DAY)
}

// LookupEnvPairs 根据所有 key 查找环境变量，并通过 map 保存，如果中途有 error 不会直接中断，而是继续
func LookupEnvPairs(envPairs *map[string]string) {
	for envKey := range *envPairs {
		env, ok := os.LookupEnv(envKey)
		if !ok {
			(*envPairs)[envKey] = ""
			continue
		}
		(*envPairs)[envKey] = env
	}
	return
}

// LookupEnv 根据所有 key 查找环境变量，并通过 map 保存，如果中途有 error 不会直接中断，而是继续
func LookupEnv(envKey string) (string, error) {
	env, ok := os.LookupEnv(envKey)
	if !ok {
		return "", errorx.ErrEnvLookupFailed
	}
	return env, nil
}

func StructToJsonStr(s any) string {
	if s == nil {
		return "{}"
	}
	res, err := json.Marshal(s)
	if err != nil {
		return "{}"
	}
	return string(res)
}
