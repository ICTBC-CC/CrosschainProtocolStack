// coding:utf-8
// 转发层
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type TRANSFER struct{}

func (t *TRANSFER) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *TRANSFER) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *TRANSFER) InvokeContract(method string) pb.Response {
	args := sdk.Instance.GetArgs()
	fromchain := args[common.KEY_FROM_CHAIN]
	tochain := args[common.KEY_TO_CHAIN]
	state := args[common.KEY_TRANSFER_STATE]
	// 是整个协议的数据包
	data := args[common.KEY_TRANSFER_DATA]

	// according method segment to select contract functions
	switch method {
	case common.FUNC_SEND:
		// 发送
		return t.send(fromchain, tochain, data)
	case common.FUNC_RECEIVE:
		// 收到消息
		return t.receive(fromchain, tochain, state)
	case common.FUNC_RESPONSE:
		// 响应peer的消息
		return t.response(data)
	default:
		return sdk.Error("invalid function")
	}
}
