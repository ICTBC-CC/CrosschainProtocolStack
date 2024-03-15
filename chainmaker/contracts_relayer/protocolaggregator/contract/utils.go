// coding:utf-8
// 一些工具函数
package contract

import (
	"CPS/chainmaker/contracts/common"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 抛出警告事件
// @param data 警告信息
func (t *ProtocolAggregator) emit_warning(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_WARNING, data)
}

// 抛出通知事件
// @param data 通知事件
func (t *ProtocolAggregator) emit_info(data []string) {
	sdk.Instance.EmitEvent(common.EVENT_INFO, data)
}
