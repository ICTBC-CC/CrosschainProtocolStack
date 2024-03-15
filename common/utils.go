// coding:utf-8
// 当前package里面的辅助函数
package common

import "fmt"

// 获取消息类型的string字符串
func GetMsgTypeString(msg_type MSG_TYPE) string {
	switch msg_type {
	case MSG_ERROR:
		return "MSG_ERROR"
	case MSG_CROSS_CHAIN:
		return "MSG_CROSS_CHAIN"
	case MSG_SEND:
		return "MSG_SEND"
	case MSG_RECEIVE:
		return "MSG_RECEIVE"
	case MSG_INFO:
		return "MSG_INFO"
	case MSG_WARNING:
		return "MSG_WARNING"
	}
	return string(msg_type)
}

// 格式化输出MSG
// @param msg 订阅解析后的消息
// @return string 消息字符串
func GetMsgInfo(msg *SubscribeMsg) string {
	info := ""
	info += "MSG_TYPE:" + GetMsgTypeString(msg.Type) + ","
	info += "ChainName:" + msg.ChainName + ","
	info += "UserName:" + msg.UserName + ","
	info += "ContractName:" + msg.ContractName + ","
	info += "swap_resource:" + fmt.Sprintf("%+v", msg.AtomicSwap) + ","
	info += "AppFrom:" + string(msg.AppFrom) + ","
	info += "AppFunc:" + string(msg.AppFunc) + ","
	info += "AppData:" + string(msg.AppData) + ","
	info += "TransactionID:" + string(msg.TransactionID) + ","
	info += "TransactionState:" + string(msg.TransactionState) + ","
	info += "TransactionData:" + string(msg.TransactionData) + ","
	info += "TransactionProtocol:" + msg.TransactionProtocol + ","
	info += "VerifyType:" + string(msg.VerifyType) + ","
	info += "VerifyData:" + string(msg.VerifyData) + ","
	info += "FromChain:" + string(msg.FromChain) + ","
	info += "ToChain:" + string(msg.ToChain) + ","
	info += "TransferState:" + string(msg.TransferState) + ","

	return info
}
