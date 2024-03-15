// coding:utf-8
// 对接收消息进行响应
package contract

import (
	"CPS/chainmaker/contracts/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 对接收消息进行响应
func (t *ProtocolAggregator) responseTransfer(msg *common.ProtocolMsg) error {
	// 序列化
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// 调用转发层
	kvs := map[string][]byte{
		common.KEY_TRANSFER_DATA: data,
	}
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSFER,
		common.FUNC_RESPONSE,
		kvs,
	)

	// 调用失败回滚事务层
	if resp.Status != sdk.OK {
		// ! 回滚事务层
		sdk.Instance.CallContract(
			common.LAYER_TRANSACTION,
			common.FUNC_REVERT,
			sdk.Instance.GetArgs(),
		)
		return fmt.Errorf("call transfer response error:%s", resp.Message)
	}
	return nil
}
