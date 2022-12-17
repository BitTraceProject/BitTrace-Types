package common

import "github.com/BitTraceProject/BitTrace-Types/pkg/constants"

func ExecuteFunctionByRetry(f func() error) error {
	var err error
	for i := 0; i < constants.RETRY_COUNT; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}
