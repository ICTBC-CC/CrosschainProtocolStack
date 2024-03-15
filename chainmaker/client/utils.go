// coding:utf-8
// 客户端内部工具函数
package client

import (
	cps_common "CPS/common"
	"CPS/utils"
	"encoding/json"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 转换监听的消息为MSG格式
func event2Subevent(
	event *common.ContractEventInfo,
	chainname, eventname, contractname string,
	msg_type cps_common.MSG_TYPE,
) *cps_common.SubscribeMsg {
	data, err := json.Marshal(event.EventData)
	if err != nil {
		info := utils.InfoError(err)
		data = []byte(info)
		msg_type = cps_common.MSG_ERROR
	}
	msg := utils.NewSucscribeMsg(msg_type, chainname, "", contractname, eventname, data)
	return msg
}
