// coding:utf-8
// 转发层receive
package contract

import (
	"CPS/chainmaker/contracts/common"
	"bytes"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层接收到消息
// @param fromchain 消息的源链
// @param tochain 消息的目的链
// @param state 消息的状态
func (t *TRANSFER) receive(fromchain, tochain, state []byte) pb.Response {
	// 转发层接收消息不做任何处理,仅判断状态
	if !bytes.Equal(state, common.TRANSFER_STATE_SEND) &&
		!bytes.Equal(state, common.TRANSFER_STATE_RESPONSE) {
		return sdk.Error("invalid transfer state")
	}
	sdk.Instance.EmitEvent(common.EVENT_TRANSFER_RECEIVE, []string{string(fromchain), string(tochain)})
	return sdk.Success(nil)
}
