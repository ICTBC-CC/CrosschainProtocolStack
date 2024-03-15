// coding:utf-8
// 事务层回滚
package contract

import (
	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// 转发层回滚
func (t *TRANSACTION) revert() pb.Response {
	_, err := t.getArgs()
	if err != nil {
		t.emit_warning([]string{"get args error:", err.Error()})
		return sdk.Error("get args error:" + err.Error())
	}
	// 抛出数据
	t.emit_info([]string{"revert transaction"})
	return sdk.Success(nil)
}
