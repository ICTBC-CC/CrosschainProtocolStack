// coding:utf-8
// 转发层发出一条响应消息给peer
package contract

import (
	"CPS/chainmaker/contracts/common"
	"bytes"
	"encoding/json"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层响应给peer
// @return response
func (t *TRANSFER) response(data []byte) pb.Response {
	// 本地处理完成,回复给对端链
	// 1. 交换源链和目的链
	// 2. 改变消息状态为响应
	msg := &common.ProtocolMsg{}
	// 解码数据包
	if err := json.Unmarshal(data, msg); err != nil {
		return sdk.Error(err.Error())
	}
	// 判断ack状态
	if bytes.Equal(msg.TransferState, common.TRANSFER_STATE_SEND) {
		// 判断ACK状态, 如果状态为send就需要改为receive,并交换源与目的链
		msg.TransferState = common.TRANSFER_STATE_RESPONSE
		tmp := msg.FromChain
		msg.FromChain = msg.ToChain
		msg.ToChain = tmp
		// 编码
		data, err := json.Marshal(msg)
		if err != nil {
			return sdk.Error(err.Error())
		}
		// 通知网关
		if err := t.sendToRelayer(msg.FromChain, msg.ToChain, data); err != nil {
			return sdk.Error(err.Error())
		}
		sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"response to peer"})
		return sdk.Success(nil)
	} else if bytes.Equal(msg.TransferState, common.TRANSFER_STATE_RESPONSE) {
		// 判断ack状态, 如果状态为receive就表示是自己发出的, 表示事务完成,不做响应
		sdk.Instance.EmitEvent(
			common.EVENT_INFO,
			[]string{
				"not deal",
				string(msg.FromChain),
				string(msg.ToChain),
			},
		)
		return sdk.Success(nil)
	}
	// 其他ack状态表示错误
	return sdk.Error("invalid msg ack " + string(msg.TransferState))
}
