// coding:utf-8
// 发送给中继链
package contract

import (
	"CPS/chainmaker/contracts/common"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 发送链上情况给链下网关
// @param fromchain 消息的源链
// @param tochain 消息的目的链
// @param data 抛出的数据
func (t *TRANSFER) sendToRelayer(fromchain, tochain, data []byte) error {
	// ! 目前就用事件的形式实现转发层的通知
	// 源链, 目的链, 整个消息的字节码
	sdk.Instance.EmitEvent(
		common.EVENT_TRANSFER_SEND,
		[]string{
			string(fromchain),
			string(tochain),
			string(data),
		},
	)
	return nil
}
