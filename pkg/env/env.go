package env

import (
	"os"

	"github.com/BitTraceProject/BitTrace-Types/pkg/errorx"
)

// LookupEnvPairs 根据所有 key 查找环境变量，并通过 map 保存，如果中途有 error 会直接中断
func LookupEnvPairs(envPairs *map[string]string) error {
	for keyStr := range *envPairs {
		env, ok := os.LookupEnv(keyStr)
		if !ok {
			return errorx.ErrEnvLookupFailed
		}
		(*envPairs)[keyStr] = env
	}
	return nil
}
