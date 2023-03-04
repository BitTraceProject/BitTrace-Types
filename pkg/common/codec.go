package common

import "encoding/base64"

// Base64StdDecode 自动解码传入的 base64 字节数组
func Base64StdDecode(src []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	dst = dst[:n]
	return dst, nil
}
