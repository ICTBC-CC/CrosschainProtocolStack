// coding:utf-8
// 执行事务层receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用事务层receive
// @param msg 消息
// @return error err
func (t *ProtocolAggregator) receiveTransaction(msg *common.ProtocolMsg) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal msg error:" + err.Error())
	}
	kvs := map[string][]byte{cps_common.KEY_MSG: msgBytes}
	// 调用事务层
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSACTION,
		common.FUNC_RECEIVE,
		kvs,
	)
	// 转发层调用失败就返回错误
	if resp.Status != sdk.OK {
		// ! 好好考虑一下事务层receive失败是否需要回滚
		return fmt.Errorf("call transaction receive error:%s", resp.Message)
	}
	return nil
}
