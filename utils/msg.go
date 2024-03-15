// coding:utf-8
// 消息工具函数
package utils

import cps_common "CPS/common"

// 制作新的subscribe msg
// @param msg_type 消息类型
// @param chainname 链名
// @param username 用户名
// @param contractname 合约名
// @param eventname 事件名
// @param data 事件数据
// @return msg msg
func NewSucscribeMsg(
	msg_type cps_common.MSG_TYPE,
	chainname, username, contractname, eventname string,
	data []byte,
) *cps_common.SubscribeMsg {
	msg := &cps_common.SubscribeMsg{
		Type:         msg_type,
		ChainName:    chainname,
		UserName:     username,
		ContractName: contractname,
		EventName:    eventname,
		ChainData:    data,
	}
	return msg
}
