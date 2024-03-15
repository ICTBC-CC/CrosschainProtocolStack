// coding:utf-8
// 调用验证层执行发送
package contract

import (
	"CPS/chainmaker/contracts/common"
	"fmt"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 调用验证层
// @param error err
func (t *ProtocolAggregator) sendVerify() error {
	// 调用验证层
	resp := sdk.Instance.CallContract(
		common.LAYER_VERIFY,
		common.FUNC_SEND,
		sdk.Instance.GetArgs(),
	)

	// 验证层执行失败就回滚
	if resp.Status != sdk.OK {
		// ! 回滚事务层
		sdk.Instance.CallContract(
			common.LAYER_TRANSACTION,
			common.FUNC_REVERT,
			sdk.Instance.GetArgs(),
		)
		return fmt.Errorf("call verify send error:%s", resp.Message)
	}
	return nil
}
