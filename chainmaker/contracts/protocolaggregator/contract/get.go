// coding:utf-8
// 获取参数
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 获取参数
// @return msg msg
// @return error err
func (t *ProtocolAggregator) getArgs() (*common.ProtocolMsg, error) {
	// 获取msg
	msgByte := sdk.Instance.GetArgs()[cps_common.KEY_MSG]
	if len(msgByte) == 0 {
		t.emit_warning([]string{"get args error", "invalid length of msg byte"})
		return nil, fmt.Errorf("invalid length of msg byte")
	}

	var msg common.ProtocolMsg
	if err := json.Unmarshal(msgByte, &msg); err != nil {
		t.emit_warning([]string{"get args unmarshal msg error", err.Error()})
		return nil, fmt.Errorf("unmarshal msg error:" + err.Error())
	}

	return &msg, nil
}
