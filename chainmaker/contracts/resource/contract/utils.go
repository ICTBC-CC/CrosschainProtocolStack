// coding:utf-8
// 工具
package contract

import (
	"CPS/chainmaker/contracts/common"
	cps_common "CPS/common"
	"encoding/json"
	"fmt"
	"strings"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 制作数据库的key
// @param app 应用名
// @param name 资源名
// @return string 制作的key
// @return error err
func (r *Resource) makeKey(app, name string) (string, error) {
	if len(app) == 0 || len(name) == 0 {
		return "", fmt.Errorf("invalid length of app or name")
	}
	if strings.Contains(app, "_") || strings.Contains(name, "_") {
		return "", fmt.Errorf("can not contains '_'")
	}
	return app + "_" + name, nil
}

// *************************** 其他 ***************************

// 抛出警告事件
// @param data 警告信息
func (t *Resource) emit_warning(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_WARNING, data)
}

// 抛出通知事件
// @param data 通知事件
func (t *Resource) emit_info(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_INFO, data)
}

// 获取参数
// @return msg msg
// @return error err
func (t *Resource) getArgs() (*common.ProtocolMsg, error) {
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
