// coding:utf-8
// 发送消息到对端链的示例
package example

import (
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 发送一条消息到目标链
// @param client 客户端
// @param msg 发送的字符串消息
// @param destChain 目标链
// @return error err
func Testapp(client cps_common.Client, msg, destChain string) error {
	// 转移的资源
	mp := map[string]int{
		"name":  1, // 转移name资源1个到对端链的app
		"name2": 2, // 转移name资源1个到对端链的app
	}
	mpBytes, err := json.Marshal(mp)
	if err != nil {
		info := utils.InfoError(err)
		panic(info)
	}

	// 调用testapp
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte("test1")},
		{Key: "data", Value: []byte("this is data")},
		// 指定验证类型
		{Key: cps_common.KEY_VERIFY_TYPE, Value: []byte(cps_common.VERIFY_TYPE_HAPPY)},
		// 指定使用的资源
		{Key: cps_common.KEY_RESOURCE_USE_FROM, Value: mpBytes},
		{Key: cps_common.KEY_RESOURCE_USE_TO, Value: mpBytes},
	}
	utils.InfoTips("start to invoke contract")
	if _, err := client.InvokeContract("testapp", "", kvs); err != nil {
		panic(err)
	}
	return nil
}
