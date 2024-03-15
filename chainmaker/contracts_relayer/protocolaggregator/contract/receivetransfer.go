// coding:utf-8
// 执行转发层receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用转发层receive
// @param error err
func (t *ProtocolAggregator) receiveTransfer(fromchain, tochain []byte) error {
	kvs := map[string][]byte{
		common.KEY_FROM_CHAIN: fromchain,
		common.KEY_TO_CHAIN:   tochain,
	}
	// 调用转发层
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSFER,
		common.FUNC_RECEIVE,
		kvs,
	)
	// 转发层调用失败就返回错误
	if resp.Status != sdk.OK {
		return fmt.Errorf("call transfer receive error:%s", resp.Message)
	}
	return nil
}
