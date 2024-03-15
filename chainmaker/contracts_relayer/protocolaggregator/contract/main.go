// coding:utf-8
// 协议聚合器
package contract

import (
	"CPS/chainmaker/contracts/common"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type ProtocolAggregator struct{}

func (t *ProtocolAggregator) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (t *ProtocolAggregator) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (t *ProtocolAggregator) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	// ! 注意, 中继链只有receive函数
	case common.FUNC_RECEIVE:
		// 从网关收到一个消息
		return t.receive()
	default:
		return sdk.Error("protocal aggregator invalid function:" + method)
	}
}
