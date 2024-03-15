// coding:utf-8
// 哈希相关
package utils

import "crypto/sha256"

// 计算sha256哈希
// @param abData 字节数组
// @return string 计算的哈希
func GetSha256Hash(abData []byte) []byte {
	hasher := sha256.New()
	hasher.Write(abData)
	return hasher.Sum(nil)
}
