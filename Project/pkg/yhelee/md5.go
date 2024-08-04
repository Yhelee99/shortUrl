package yhelee

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5Value 获取md5的值
func GetMd5Value(v []byte) string {
	h := md5.New()
	// h.Write() 将实际的数据写入哈希计算器
	h.Write(v)
	// h.Sum() 传空就是获取当前的哈希结果，如果传入参数b，会在计算出的哈希值的头部添加上b
	return hex.EncodeToString(h.Sum(nil))
}
