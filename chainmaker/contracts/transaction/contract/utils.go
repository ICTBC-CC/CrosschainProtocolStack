// coding:utf-8
// 工具函数
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 抛出警告事件
// @param data 警告信息
func (t *TRANSACTION) emit_warning(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_WARNING, data)
}

// 抛出通知事件
// @param data 通知事件
func (t *TRANSACTION) emit_info(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_INFO, data)
}

// 获取参数
// @return msg msg
// @return error err
func (t *TRANSACTION) getArgs() (*common.ProtocolMsg, error) {
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
