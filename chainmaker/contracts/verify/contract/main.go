// coding:utf-8
// 验证层
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type VERIFY struct{}

func (t *VERIFY) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *VERIFY) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *VERIFY) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	case common.FUNC_SEND:
		// 发送
		return t.send()
	case common.FUNC_RECEIVE:
		// 收到消息
		return t.receive()
	case common.FUNC_VERIFY_ADD_NOTORY:
		// 注册notary
		return t.addNotary()
	default:
		return sdk.Error("invalid function")
	}
}
