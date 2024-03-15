// coding:utf-8
// 执行转发层的send
package contract

import (
	"CPS/chainmaker/contracts/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用转发层
// @param fromchain 源链
// @param tochain 目的链
// @param error err
func (t *ProtocolAggregator) sendTransfer(protoMsg *common.ProtocolMsg) error {
	data, err := json.Marshal(protoMsg)
	kvs := map[string][]byte{
		common.KEY_FROM_CHAIN:    protoMsg.FromChain,
		common.KEY_TO_CHAIN:      protoMsg.ToChain,
		common.KEY_TRANSFER_DATA: data,
	}
	// 调用转发层
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSFER,
		common.FUNC_SEND,
		kvs,
	)

	// 转发层执行失败 或 json.marshal失败 就回滚
	if resp.Status != sdk.OK || err != nil {
		// ! 回滚事务层
		sdk.Instance.CallContract(
			common.LAYER_TRANSACTION,
			common.FUNC_REVERT,
			sdk.Instance.GetArgs(),
		)

		if err == nil {
			return fmt.Errorf("call transfer send marshal error:%s", err)
		} else {
			return fmt.Errorf("call transfer send error:%s", resp.Message)
		}
	}
	return nil
}
