// coding:utf-8
// 转发层发送
package contract

import (
	"CPS/chainmaker/contracts/common"
	"encoding/json"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层发送
// @param fromchain 源链
// @param tochain 目的链
// @param data 序列化的全部消息
func (t *TRANSFER) send(fromchain, tochain, data []byte) pb.Response {
	msg := &common.ProtocolMsg{}
	// 解码数据包
	if err := json.Unmarshal(data, msg); err != nil {
		return sdk.Error(err.Error())
	}
	// 改变消息状态为send状态
	if msg.TransferState != nil {
		// 初始发过来的消息状态必定为[]byte{}空
		return sdk.Error("invalid transfer state")
	}

	msg.TransferState = common.TRANSFER_STATE_SEND
	msg.FromChain = fromchain
	msg.ToChain = tochain

	// 编码
	data, err := json.Marshal(msg)
	if err != nil {
		return sdk.Error(err.Error())
	}

	// 通知给中继网关
	if err := t.sendToRelayer(fromchain, tochain, data); err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(nil)
}
