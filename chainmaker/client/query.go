// coding:utf-8
// 查询合约
package client

import (
	"CPS/utils"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 查询合约信息
// @param client: 客户端
// @param contractName: 合约名称
// @param method: 调用或者查询，这里一般是设置成invoke_contract
// @param txId: 可以置空
// @param kvs: 参数列表
// @param withSyncResult: 是否同步获取结果
func (c *Client) QueryContract(contractName string, kvs []*common.KeyValuePair) (*common.TxResponse, error) {

	resp, err := c.ChainmakerSDK.QueryContract(contractName, "invoke_contract", kvs, -1)
	if err != nil {
		return resp, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return resp, fmt.Errorf("invoke %s contract failed. [code]:%d\t[msg]:%s", contractName, resp.Code, resp.Message)
	}
	// fmt.Printf("query %s contract success: [contractResult]:%s\n", contractName, resp.ContractResult)
	utils.Info("query " + contractName + " contract success: [contractResult]:" + resp.ContractResult.String() + " [msg]:" + resp.Message)
	return resp, nil
}
