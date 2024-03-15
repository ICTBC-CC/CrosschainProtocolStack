// coding:utf-8
// 调用合约
package client

import (
	"CPS/utils"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 调用合约方法
// @param contractName 合约名称
// @param method 方法名
// @param kvs 参数列表
// @return any result
// @return error err
func (c *Client) InvokeContract(
	contractName string,
	method string,
	args ...any,
) (any, error) {
	// 解析参数
	kvs, err := parseArgs(method, args...)
	if err != nil {
		info := utils.InfoError(err)
		return nil, fmt.Errorf(info)
	}

	// 调用合约
	resp, err := c.ChainmakerSDK.InvokeContract(contractName, "invoke_contract", "", kvs, -1, true)
	if err != nil {
		info := utils.InfoError(err)
		return resp, fmt.Errorf(info)
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		info := utils.InfoTips("call contract failed", resp.ContractResult)
		return resp, fmt.Errorf(info)
	}

	return resp, nil
}

// 解析args
func parseArgs(method string, args ...any) ([]*common.KeyValuePair, error) {
	kvs := []*common.KeyValuePair{}
	if len(args) == 0 {
		// 没有参数,就只加入method
		kvs = append(kvs, &common.KeyValuePair{
			Key:   "method",
			Value: []byte(method),
		})
		return kvs, nil
	}

	// 有参数,首先判断传入的类型是不是格式化的参数
	var ok bool
	kvs, ok = args[0].([]*common.KeyValuePair)
	if !ok {
		// 不ok表示不是传入的参数对, 就用data字段
		// 转化为字节数组
		data, ok := args[0].([]byte)
		if !ok {
			// 不ok表示参数错误
			return kvs, fmt.Errorf("chainmaker convert args not ok")
		}

		// 制作参数
		kvs = append(kvs, &common.KeyValuePair{
			Key:   "data",
			Value: data,
		})
	}

	// 判断method参数是否存在
	has_method := false
	for _, kv := range kvs {
		if kv.Key == "method" {
			has_method = true
			break
		}
	}
	// 加入方法参数
	if !has_method {
		kvs = append(kvs, &common.KeyValuePair{
			Key:   "method",
			Value: []byte(method),
		})
	}

	return kvs, nil
}
