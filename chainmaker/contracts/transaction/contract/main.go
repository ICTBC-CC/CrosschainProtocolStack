// coding:utf-8
// 事务层
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type TRANSACTION struct{}

func (t *TRANSACTION) InitContract() pb.Response {
	// 初始化初始id
	currID := common.NIL_TRANSACTION_ID + 1

	// 初始化最新id
	if err := t.updateNesestID(currID); err != nil {
		return sdk.Error("update newest transaction id error:" + err.Error())
	}

	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *TRANSACTION) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *TRANSACTION) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	case common.FUNC_SEND:
		// 发送
		return t.send()
	case common.FUNC_RECEIVE:
		// 收到消息
		return t.receive()
	case common.FUNC_REVERT:
		// 回滚
		return t.revert()
	default:
		return sdk.Error("invalid function")
	}
}
