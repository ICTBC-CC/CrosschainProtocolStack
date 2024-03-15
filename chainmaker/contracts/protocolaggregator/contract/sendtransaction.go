// coding:utf-8
// 调用事务层send
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用事务层send
// @param msg 协议消息
// @return []byte transaction层返回的消息
// @return error err
func (t *ProtocolAggregator) sendTransaction(msg *common.ProtocolMsg) ([]byte, error) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal msg error:" + err.Error())
	}

	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}
	// 调用事务层,返回事务id
	resp := sdk.Instance.CallContract(
		common.LAYER_TRANSACTION,
		common.FUNC_SEND,
		kvs,
	)
	// !事务层失败后是没有回滚的,因为本来就不会改变状态
	if resp.Status != sdk.OK {
		return nil, fmt.Errorf("call transaction send error:%s", resp.Message)
	}

	// 返回事务id
	return resp.Payload, nil
}
