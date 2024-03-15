// coding:utf-8
// 执行应用层的send
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用应用层send
// @param msg 协议消息
// @return []byte 应用层数据
// @return error err
func (t *ProtocolAggregator) sendApp(msg *common.ProtocolMsg) ([]byte, error) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal msg error:" + err.Error())
	}

	kvs := map[string][]byte{cps_common.KEY_MSG: msgByte}
	// 调用app,返回跨链数据
	resp := sdk.Instance.CallContract(
		string(msg.FromApp),
		common.FUNC_SEND,
		kvs,
	)

	// app层失败后是没有回滚的,因为本来就不会改变状态
	if resp.Status != sdk.OK {
		return nil, fmt.Errorf("call app %s send error:%s", msg.FromApp, resp.Message)
	}

	// 返回应用层数据
	return resp.Payload, nil
}
