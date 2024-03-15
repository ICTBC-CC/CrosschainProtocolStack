// coding:utf-8
// 执行app层的receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用应用层receive
// @param msg 协议消息
// @return []byte 应用层数据
// @return error err
func (t *ProtocolAggregator) receiveApp(msg *common.ProtocolMsg) ([]byte, error) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal msg error:" + err.Error())
	}

	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}
	// 调用app,返回跨链数据
	resp := sdk.Instance.CallContract(
		string(msg.FromApp),
		common.FUNC_RECEIVE,
		kvs,
	)

	// app层失败后回滚事务层
	if resp.Status != sdk.OK {
		// 验证层执行失败就回滚
		if resp.Status != sdk.OK {
			// ! 回滚事务层
			sdk.Instance.CallContract(
				common.LAYER_TRANSACTION,
				common.FUNC_REVERT,
				sdk.Instance.GetArgs(),
			)
			return nil, fmt.Errorf("call app receive error:%s", resp.Message)
		}
		return nil, fmt.Errorf("call app send error:%s", resp.Message)
	}

	// 返回应用层返回的数据
	return resp.Payload, nil
}
