// coding:utf-8
// 验证层发送
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层发送
func (t *VERIFY) send() pb.Response {
	// 抛出数据给链下网关
	sdk.Instance.EmitEvent(common.EVENT_INFO, []string{"verify send"})
	return sdk.Success(nil)
}
